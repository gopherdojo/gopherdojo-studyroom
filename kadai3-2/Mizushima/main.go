package main

import (
	"context"
	"fmt"
	"log"
	"net/http/httputil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	download "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/download"
	getheader "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/getheader"
	listen "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/listen"
	options "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/options"
	"github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/request"
)

// for test
// go run main.go https://4.bp.blogspot.com/-2-Ny23XgrF0/Ws69gszw2jI/AAAAAAABLdU/unbzWD_U8foWBwPKWQdGP1vEDoQoYjgZwCLcBGAs/s1600/top_banner.jpg -o kadai3-2
// go run main.go http://i.imgur.com/z4d4kWk.jpg -o .
// go run main.go https://misc.laboradian.com/test/003/ -o .

func main() {
	opts, urls, err := options.ParseOptions(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if opts.Procs == 0 {
		opts.Procs = uint(runtime.NumCPU())
	}

	fmt.Printf("opts:%#v\n", opts)
	fmt.Println(urls)
	
	if len(opts.Output) > 0 && opts.Output[len(opts.Output)-1] != '/' {
		opts.Output += "/"
	}

	for i, url := range urls {

		resp, err := request.Request("HEAD", url, "", "")
		if err != nil {
			log.Fatalf("err: %s\n", err)
		}
		
		fmt.Println("response:")
		bytes, err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Fatalf("err: httputil.DumpResponse: %s\n", err)
		}
		fmt.Println(string(bytes))

		fileSize, err := getheader.GetSize(resp)
		if err != nil {
			log.Fatalf("err: %s\n", err)
		}
		resp.Body.Close()

		partial := fileSize / opts.Procs

		out, err := os.OpenFile(opts.Output + filepath.Base(url), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			log.Fatalf("err: os.Create: %s\n", err)
		}

		// make a temporary directory
		tmpDirName := opts.Output + strconv.Itoa(i)
		err = os.Mkdir(tmpDirName, 0777)
		if err != nil {
			out.Close()
			if err2 := os.Remove(opts.Output + filepath.Base(url)); err2 != nil {
				log.Fatalf("err: os.Mkdir: %s\nerr: os.Remove: %s\n", err, err2)
			}
			log.Fatalf("err: os.Mkdir: %s\n", err)
		}

		// ctx, cancel := context.WithTimeout(context.Background(),time.Duration(opts.Tm)*time.Minute)
		ctx := context.Background()
		clean := func() { 
			out.Close()
			// delete the tmporary directory
			if err := os.RemoveAll(tmpDirName); err != nil {
				log.Fatalf("err: RemoveAll: %s\n", err)
			}
			if err := os.Remove(opts.Output + filepath.Base(url)); err != nil {
				log.Fatalf("err: os.Remove: %s\n", err)
			}
		}
		ctx, cancel := listen.Listen(ctx, os.Stdout, clean)

		var isPara bool = true
		accept, err := getheader.ResHeader(os.Stdout, resp, "Accept-Ranges")
		if err != nil && err.Error() == "cannot find Accept-Ranges header" {
			isPara = false  
		} else if err != nil {
			clean()
			log.Fatalf("err: getheader.ResHeader: %s\n", err)
		} else if accept[0] != "bytes" {
			isPara = false
			continue;
		}
		
		err = download.Downloader(url, out, fileSize, partial, opts.Procs, isPara, tmpDirName, ctx)
		if err != nil {
			log.Fatalf("err: %s\n", err)
		}

		fmt.Println("download complete")

		err = MergeFiles(tmpDirName, opts.Procs, fileSize, out)
		if err != nil {
			log.Fatalf("err: MergeFiles: %s\n", err)
		}

		// delete the tmporary directory
		if err := os.RemoveAll(tmpDirName); err != nil {
			log.Fatalf("err: RemoveAll: %s\n", err)
		}

		cancel()
		out.Close()
	}
}

func MergeFiles(tmpDirName string, procs, fileSize uint, output *os.File) error {
	for i := uint(0); i < procs; i++ {

		body, err := os.ReadFile(tmpDirName+"/"+strconv.Itoa(int(i)))
		if err != nil {
			return err
		}

		fmt.Fprint(output, string(body))
		fmt.Printf("target file: %s, len=%d written\n", output.Name(), len(string(body)))
	}
	return nil
}
