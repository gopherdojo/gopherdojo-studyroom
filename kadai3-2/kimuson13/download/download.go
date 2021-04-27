package download

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

// Donwnloader struct
type Downloader struct {
	parallel int
	timeout  int
	filename string
	url      string
}

// Rnage struct
type Range struct {
	low    int
	high   int
	number int
}

// New for download package
func New(opts *Options) *Downloader {
	return &Downloader{
		parallel: opts.Parallel,
		timeout:  opts.Timeout,
		filename: opts.Filename,
		url:      opts.URL,
	}
}

func CreateTempdir() error {
	if err := os.Mkdir("tempdir", 0755); err != nil {
		return err
	}

	return nil
}

func DeleteTempdir() {
	err := os.RemoveAll("tempdir")
	if err != nil {
		log.Println("can't remove the tempdir", err)
	}
}

// Run excecute method in download package
func (d *Downloader) Run(ctx context.Context) error {
	if err := d.Validation(); err != nil {
		return err
	}

	if err := CreateTempdir(); err != nil {
		return err
	}
	defer DeleteTempdir()

	contentLength, err := d.CheckContentLength(ctx)
	if err != nil {
		return err
	}

	if err := d.Download(contentLength, ctx); err != nil {
		return err
	}

	if err := d.MergeFile(d.parallel, contentLength); err != nil {
		return err
	}

	return nil
}

//Preparate method define the variables to Donwload
func (d *Downloader) Validation() error {
	if d.parallel < 1 {
		return errors.New("the parallel number needs to be bigger than 1")
	}

	if d.timeout < 1 {
		return errors.New("the timeout needs to be bigeer than 1")
	}

	return nil
}

// CheckContentLength method gets the Content-Length want to download
func (d *Downloader) CheckContentLength(ctx context.Context) (int, error) {
	if _, err := fmt.Fprintf(os.Stdout, "Start HEAD request to check Content-Length\n"); err != nil {
		return 0, err
	}

	req, err := http.NewRequest("HEAD", d.url, nil)
	if err != nil {
		return 0, err
	}

	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	acceptRange := res.Header.Get("Accept-Ranges")
	if _, err := fmt.Fprintf(os.Stdout, "got: Accept-Ranges: %s\n", acceptRange); err != nil {
		return 0, err
	}

	if acceptRange == "" {
		return 0, errors.New("Accept-Range is not bytes")
	}

	if acceptRange != "bytes" {
		return 0, errors.New("it is not content")
	}

	contentLength := int(res.ContentLength)
	if _, err := fmt.Fprintf(os.Stdout, "got: Content-Length: %v\n", contentLength); err != nil {
		return 0, err
	}

	if contentLength < 1 {
		return 0, errors.New("it does not have Content-Length")
	}

	return contentLength, nil
}

// Download method does split-download with goroutine
func (d *Downloader) Download(contentLength int, ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(d.timeout)*time.Second)
	defer cancel()

	parallel := d.parallel
	split := contentLength / parallel
	grp, ctx := errgroup.WithContext(ctx)
	for i := 0; i < parallel; i++ {
		r := MakeRange(i, split, parallel, contentLength)
		tempfile := fmt.Sprintf("tempdir/tempfile.%d.%d", parallel, r.number)
		filename, err := CreateTempfile(tempfile)
		if err != nil {
			return err
		}

		grp.Go(func() error {
			err := Requests(r, d.url, filename)
			return err
		})
	}

	if err := grp.Wait(); err != nil {
		return err
	}

	return nil
}

func CreateTempfile(name string) (string, error) {
	file, err := os.Create(name)
	if err != nil {
		return "", err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Println("can't close the "+name, err)
		}
	}()

	return file.Name(), nil
}

// MakeRange function distributes Content-Length for split-download
func MakeRange(i, split, parallel, contentLength int) Range {
	low := split * i
	high := low + split - 1
	if i == parallel-1 {
		high = contentLength
	}

	return Range{
		low:    low,
		high:   high,
		number: i,
	}
}

// Requests function sends GET request
func Requests(r Range, url, filename string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", r.low, r.high))
	if _, err := fmt.Fprintf(os.Stdout, "start GET request: bytes=%d-%d\n", r.low, r.high); err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.New("error is here")
	}

	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Println("the response body can't close", err)
		}
	}()

	if res.StatusCode != http.StatusPartialContent {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	output, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer func() {
		err := output.Close()
		if err != nil {
			log.Println("can't close the tempfile", err)
		}
	}()

	_, err = io.Copy(output, res.Body)
	if err != nil {
		return err
	}
	return nil
}

// MergeFile method merges tempfiles to new file
func (d *Downloader) MergeFile(parallel, contentLength int) error {
	fmt.Println("\nmerging files...")
	filename := d.filename
	fh, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer func() {
		err := fh.Close()
		if err != nil {
			log.Println("can't close the download file!", err)
		}
	}()

	for i := 0; i < parallel; i++ {
		if err := Merger(parallel, i, fh); err != nil {
			return err
		}
	}

	fmt.Println("complete parallel donwload")
	return nil
}

func Merger(parallel, i int, fh *os.File) error {
	f := fmt.Sprintf("tempdir/tempfile.%d.%d", parallel, i)
	sub, err := os.Open(f)
	if err != nil {
		return err
	}

	_, err = io.Copy(fh, sub)
	if err != nil {
		return err
	}

	defer func() {
		err := sub.Close()
		if err != nil {
			log.Println("can't close the "+f, err)
		}
	}()
	return nil
}
