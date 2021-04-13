package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

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
