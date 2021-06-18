package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	options "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/options"
)

// go run main.go https://4.bp.blogspot.com/-2-Ny23XgrF0/Ws69gszw2jI/AAAAAAABLdU/unbzWD_U8foWBwPKWQdGP1vEDoQoYjgZwCLcBGAs/s1600/top_banner.jpg


func main() {
	opts, url, err := options.ParseOptions(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("opts:%#v\n", opts)
	fmt.Println(url)

	if err := DownloadFile(opts.Output, url); err != nil {
		log.Fatal(err)
	}
}

func DownloadFile(path string, urls []string) error {
	
	for _, url := range urls {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		fmt.Printf("%#v\n", req)

		req.Header.Set("Range", "byte=0-499")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if path[len(path)-1] != '/' {
			path += "/"
		}
		
		out, err := os.Create(path+filepath.Base(url))
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return err
		}
	}
	return nil
}