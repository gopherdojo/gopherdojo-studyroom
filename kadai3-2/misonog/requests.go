package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

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

func download(filename string, dirname string, url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	filepath := fmt.Sprintf("%s/%s", dirname, filename)

	out, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer out.Close()

	io.Copy(out, res.Body)
	return nil
}
