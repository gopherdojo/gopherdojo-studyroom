package downloading

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"split_download/opt"
	"split_download/termination"
	"sync"
	"time"
)

var (
	errResponseDoesNotIncludeAcceptRangesHeader = errors.New("response does not include Accept-Ranges header")
	errValueOfAcceptRangesHeaderIsNotBytes = errors.New("the value of Accept-Ranges header is not bytes")
	errNoContent                           = errors.New("no content")
)

// Downloader has the information for the download.
type Downloader struct {
	outStream   io.Writer
	url         *url.URL
	parallelism int
	output      string
	timeout     time.Duration
}

// NewDownloader generates Downloader based on Options.
func NewDownloader(w io.Writer, opts *opt.Options) *Downloader {
	return &Downloader{
		outStream:   w,
		url:         opts.URL,
		parallelism: opts.Parallelism,
		output:      opts.Output,
		timeout:     opts.Timeout,
	}
}

// Download performs parallel download.
func (d *Downloader) Download(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, d.timeout)
	defer cancel()

	contentLength, err := d.getContentLength(ctx)
	if err != nil {
		return err
	}

	rangeHeaders := d.toRangeHeaders(contentLength)

	tempDir, err := ioutil.TempDir("", "parallel-download")
	if err != nil {
		return err
	}
	clean := func() { os.RemoveAll(tempDir) }
	defer clean()
	termination.CleanFunc(clean)

	filenames, err := d.parallelDownload(ctx, rangeHeaders, tempDir)
	if err != nil {
		return err
	}

	filename, err := d.concat(filenames, tempDir)
	if err != nil {
		return err
	}

	fmt.Fprintf(d.outStream, "rename %q to %q\n", filename, d.output)

	err = os.Rename(filename, d.output)
	if err != nil {
		return err
	}

	fmt.Fprintf(d.outStream, "completed: %q\n", d.output)

	return nil
}

// getContentLength returns the value of Content-Length received by making a HEAD request.
func (d *Downloader) getContentLength(ctx context.Context) (int, error) {
	fmt.Fprintf(d.outStream, "start HEAD request to get Content-Length\n")

	req, err := http.NewRequest("HEAD", d.url.String(), nil)
	if err != nil {
		return 0, err
	}
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	err = d.validateAcceptRangesHeader(resp)
	if err != nil {
		return 0, err
	}

	contentLength := int(resp.ContentLength)

	fmt.Fprintf(d.outStream, "got: Content-Length: %d\n", contentLength)

	if contentLength < 1 {
		return 0, errNoContent
	}

	return contentLength, nil
}

// validateAcceptRangesHeader validates the following.
// - The presence of an Accept-Ranges header
// - The value of the Accept-Ranges header is "bytes"
func (d *Downloader) validateAcceptRangesHeader(resp *http.Response) error {
	acceptRangesHeader := resp.Header.Get("Accept-Ranges")

	fmt.Fprintf(d.outStream, "got: Accept-Ranges: %s\n", acceptRangesHeader)

	if acceptRangesHeader == "" {
		return errResponseDoesNotIncludeAcceptRangesHeader
	}

	if acceptRangesHeader != "bytes" {
		return errValueOfAcceptRangesHeaderIsNotBytes
	}

	return nil
}

// toRangeHeaders converts the value of Content-Length to the value of Range header.
func (d *Downloader) toRangeHeaders(contentLength int) []string {
	parallelism := d.parallelism

	// 1 <= parallelism <= Content-Length
	if parallelism < 1 {
		parallelism = 1
	}
	if contentLength < parallelism {
		parallelism = contentLength
	}

	unitLength := contentLength / parallelism
	remainingLength := contentLength % parallelism

	rangeHeaders := make([]string, 0)

	cntr := 0
	for n := parallelism; n > 0; n-- {
		min := cntr
		max := cntr + unitLength - 1

		// Add the remaining length to the last chunk
		if n == 1 && remainingLength != 0 {
			max += remainingLength
		}

		rangeHeaders = append(rangeHeaders, fmt.Sprintf("bytes=%d-%d", min, max))

		cntr += unitLength
	}

	return rangeHeaders
}

// parallelDownload downloads in parallel for each specified rangeHeaders and saves it in the specified dir.
func (d *Downloader) parallelDownload(ctx context.Context, rangeHeaders []string, dir string) (map[int]string, error) {
	filenames := map[int]string{}

	filenameCh := make(chan map[int]string)
	errCh := make(chan error)

	for i, rangeHeader := range rangeHeaders {
		go d.partialDownloadAndSendToChannel(ctx, i, rangeHeader, filenameCh, errCh, dir)
	}

	eg, ctx := errgroup.WithContext(ctx)
	var mu sync.Mutex
	for i := 0; i < len(rangeHeaders); i++ {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case m := <-filenameCh:
				for k, v := range m {
					mu.Lock()
					filenames[k] = v
					mu.Unlock()
				}
				return nil
			case err := <-errCh:
				return err
			}
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return filenames, nil
}

// partialDownloadAndSendToChannel performs partialDownload and sends it to the appropriate channel according to the result.
func (d *Downloader) partialDownloadAndSendToChannel(ctx context.Context, i int, rangeHeader string, filenameCh chan<- map[int]string, errCh chan<- error, dir string) {
	filename, err := d.partialDownload(ctx, rangeHeader, dir)
	if err != nil {
		errCh <- err
		return
	}

	filenameCh <- map[int]string{i: filename}

	return
}

// partialDownload sends a partial request with the specified rangeHeader,
// and saves the response body in the file under the specified dir,
// and returns the filename.
func (d *Downloader) partialDownload(ctx context.Context, rangeHeader string, dir string) (string, error) {
	req, err := http.NewRequest("GET", d.url.String(), nil)
	if err != nil {
		return "", err
	}
	req = req.WithContext(ctx)

	req.Header.Set("Range", rangeHeader)

	fmt.Fprintf(d.outStream, "start GET request with header: \"Range: %s\"\n", rangeHeader)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusPartialContent {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	fp, err := os.Create(path.Join(dir, randomHexStr()))
	if err != nil {
		return "", err
	}

	_, err = io.Copy(fp, resp.Body)
	if err != nil {
		return "", err
	}

	filename := fp.Name()

	fmt.Fprintf(d.outStream, "downloaded: %q\n", filename)

	return filename, nil
}

// concat concatenates the files in order based on the mapping of the specified filenames,
// and creates the concatenated file under the specified dir,
// and returns the filename.
func (d *Downloader) concat(filenames map[int]string, dir string) (string, error) {
	fp, err := os.Create(filepath.Join(dir, randomHexStr()))
	if err != nil {
		return "", err
	}
	defer fp.Close()

	filename := fp.Name()

	fmt.Fprintf(d.outStream, "concatenate downloaded files to tempfile: %q\n", filename)

	for i := 0; i < len(filenames); i++ {
		src, err := os.Open(filenames[i])
		if err != nil {
			return "", err
		}

		_, err = io.Copy(fp, src)
		if err != nil {
			return "", err
		}
	}

	return filename, nil
}

// randomHexStr returns a random hex string of length 10.
// 10 is a length which does not duplicate enough.
func randomHexStr() string {
	b := make([]byte, 5)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", b)
}