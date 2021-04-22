package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

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
func (p *Pdownload) Check() error {
	res, err := http.Head(p.URL)
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
	p.SetFileName(filename)
	p.SetFullFileName(p.TargetDir, filename)

	p.SetFileSize(uint(res.ContentLength))

	return nil
}

// Download method distributes the task to each goroutine
func (p *Pdownload) Download() error {
	procs := uint(p.Procs)
	filesize := p.FileSize()

	// calculate split file size
	split := filesize / procs

	grp, _ := errgroup.WithContext(context.Background())

	p.Assignment(grp, procs, split)

	// wait for Assignment method
	if err := grp.Wait(); err != nil {
		return err
	}

	return nil
}

func (p Pdownload) Assignment(grp *errgroup.Group, procs, split uint) {
	filename := p.FileName()
	dirname := p.DirName()

	for i := uint(0); i < procs; i++ {
		partName := fmt.Sprintf("%s/%s.%d.%d", dirname, filename, procs, i)

		// make range
		r := p.Utils.MakeRange(i, split, procs)
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
			return p.Requests(r, filename, dirname, p.URL)
		})
	}

}

// Requests method will download the file
func (p Pdownload) Requests(r Range, filename, dirname, url string) error {
	res, err := p.MakeResponse(r, url)
	if err != nil {
		return fmt.Errorf("failed to split get requests: %d", r.woker)
	}
	defer res.Body.Close()

	partName := fmt.Sprintf("%s/%s.%d.%d", dirname, filename, p.Procs, r.woker)

	output, err := os.OpenFile(partName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 06666)
	if err != nil {
		return fmt.Errorf("failed to create %s in %s", filename, dirname)
	}
	defer output.Close()

	io.Copy(output, res.Body)

	return nil
}

// MakeResponse return *http.Respnse include context and range header
func (p Pdownload) MakeResponse(r Range, url string) (*http.Response, error) {
	// create get request
	req, err := http.NewRequest("GET", url, nil)
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
