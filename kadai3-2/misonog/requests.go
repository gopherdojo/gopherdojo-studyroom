package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/context/ctxhttp"
	"golang.org/x/sync/errgroup"
)

// Range struct for range access
type Range struct {
	low   uint
	high  uint
	woker uint
}

func isLastProc(i, procs uint) bool {
	return i == procs-1
}

// Check method check be able to range access.
func (p *Pdownload) Check(ctx context.Context, dir string) error {
	res, err := ctxhttp.Head(ctx, http.DefaultClient, p.URL)
	if err != nil {
		return err
	}

	if res.Header.Get("Accept-Ranges") != "bytes" {
		return fmt.Errorf("not supported range access: %s", p.URL)
	}

	if res.ContentLength <= 0 {
		return errors.New("invalid content length")
	}

	filename := p.Utils.FileName()
	if filename == "" {
		filename = path.Base(p.URL)
	}
	// p.SetFileName(filename)
	// p.SetFullFileName(p.TargetDir, filename)
	// p.Utils.SetDirName(dir)
	p.filename = filename
	p.setFullFileName(p.TargetDir, filename)
	p.dirname = dir
	p.filesize = uint(res.ContentLength)

	// p.SetFileSize(uint(res.ContentLength))

	return nil
}

// Download method distributes the task to each goroutine
func (p *Pdownload) Download(ctx context.Context) error {
	procs := uint(p.Procs)
	// filesize := p.FileSize()
	filesize := p.filesize

	// calculate split file size
	split := filesize / procs

	grp, ctx := errgroup.WithContext(ctx)

	p.Assignment(grp, ctx, procs, split)

	// wait for Assignment method
	if err := grp.Wait(); err != nil {
		return err
	}

	return nil
}

func (p Pdownload) Assignment(grp *errgroup.Group, ctx context.Context, procs, split uint) {
	// filename := p.FileName()
	// dirname := p.DirName()
	filename := p.filename
	dirname := p.dirname

	for i := uint(0); i < procs; i++ {
		partName := fmt.Sprintf("%s/%s.%d.%d", dirname, filename, procs, i)

		// make range
		// r := p.Utils.MakeRange(i, split, procs)
		r := p.makeRange(i, split, procs)
		if info, err := os.Stat(partName); err == nil {
			infosize := uint(info.Size())
			// check if the part is fully downloaded
			if isLastProc(i, procs) {
				if infosize == r.high-r.low {
					continue
				}
			} else if infosize == split {
				continue
			}

			// make low range from this next byte
			r.low += infosize
		}

		// execute get request
		grp.Go(func() error {
			return p.Requests(ctx, r, filename, dirname, p.URL)
		})
	}

}

// Requests method will download the file
func (p Pdownload) Requests(ctx context.Context, r Range, filename, dirname, url string) error {
	res, err := p.MakeResponse(ctx, r, url)
	if err != nil {
		// return fmt.Errorf("failed to split get requests: %d", r.woker)
		return err
	}
	defer res.Body.Close()

	partName := fmt.Sprintf("%s/%s.%d.%d", dirname, filename, p.Procs, r.woker)

	output, err := os.OpenFile(partName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 06666)
	if err != nil {
		return fmt.Errorf("failed to create %s in %s", filename, dirname)
	}
	defer output.Close()

	if _, err := io.Copy(output, res.Body); err != nil {
		return err
	}

	return nil
}

// MakeResponse return *http.Respnse include context and range header
func (p Pdownload) MakeResponse(ctx context.Context, r Range, url string) (*http.Response, error) {
	// create get request
	// req, err := http.NewRequest("GET", url, nil)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to split NewRequest for get: %d", r.woker)
	}

	// set download ranges
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", r.low, r.high))

	// set useragent
	if p.useragent != "" {
		req.Header.Set("User-Agent", p.useragent)
	}

	// set referer
	if p.referer != "" {
		req.Header.Set("Referer", p.referer)
	}

	return http.DefaultClient.Do(req)
}
