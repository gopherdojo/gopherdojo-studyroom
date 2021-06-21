package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	download "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/download"
	getheader "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/getheader"
	options "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/options"
)

// go run main.go https://4.bp.blogspot.com/-2-Ny23XgrF0/Ws69gszw2jI/AAAAAAABLdU/unbzWD_U8foWBwPKWQdGP1vEDoQoYjgZwCLcBGAs/s1600/top_banner.jpg -o kadai3-2
// go run main.go http://i.imgur.com/z4d4kWk.jpg -o .
// go run main.go https://misc.laboradian.com/test/003/ -o .

func main() {
	opts, urls, err := options.ParseOptions(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if opts.Procs == 0 {
		opts.Procs = runtime.NumCPU()
	}

	fmt.Printf("opts:%#v\n", opts)
	fmt.Println(urls)
	
	if len(opts.Output) > 0 && opts.Output[len(opts.Output)-1] != '/' {
		opts.Output += "/"
	}

	for _, url := range urls {

		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("err: %s\n", err)
		}
		defer resp.Body.Close()
		
		fileSize, err := getheader.GetSize(resp)
		if err != nil {
			log.Fatalf("err: %s\n", err)
		}
		partial := fileSize / opts.Procs

		ctx, cancel := context.WithTimeout(context.Background(),time.Duration(opts.Tm)*time.Minute)
		out, err := os.Create(opts.Output + filepath.Base(url))
		if err != nil {
			log.Fatalf("err: %s\n", err)
		}
		defer out.Close()

		accept, err := getheader.ResHeader(os.Stdout, resp, "Accept-Ranges")
		if err != nil {
			log.Fatalf("err: %s\n", err)
		} else if accept[0] != "bytes" {
			download.DownloadFile(url, out)
			continue;
		}
		
		err = pararel(url, out, fileSize, partial, opts.Procs, ctx)
		if err != nil {
			log.Fatalf("err: %s\n", err)
		}
		cancel()
	}

	
	// dump, _ := httputil.DumpResponse(resp, false)
	// fmt.Printf("response:\n%s\n", dump)

	// fmt.Printf("response status code: %s\n", resp.Status)
	// outs, _ := getheader.ResHeaderComma(os.Stdout, resp, "Accept-Ranges")
	// fmt.Println(outs == "bytes")
	// if err := download.DownloadFile(opts.Output, url); err != nil {
	// 	log.Fatal(err)
	// }
}

func pararel(url string, file *os.File, fileSize int, part int, procs int, ctx context.Context) error {
	// fmt.Println(url)
	ch := make(chan int)
	for i := 0; i < procs; i++ {
		ch <- i
		select {
		case <-ctx.Done():
			return errors.New("time limit exceeded")
		case <-ch:
			download.PDownload(url, file, fileSize, i, part, procs)
		}
	}
	return nil
}
