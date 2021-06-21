package main

import (
	"fmt"
	"log"
	"net/http/httputil"
	"os"
	"runtime"

	download "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/download"
	getheader "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/getheader"
	options "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/options"
	request "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/request"
)

// go run main.go https://4.bp.blogspot.com/-2-Ny23XgrF0/Ws69gszw2jI/AAAAAAABLdU/unbzWD_U8foWBwPKWQdGP1vEDoQoYjgZwCLcBGAs/s1600/top_banner.jpg -o kadai3-2
// go run main.go http://i.imgur.com/z4d4kWk.jpg -o .
// go run main.go https://misc.laboradian.com/test/003/ -o .

func main() {
	opts, url, err := options.ParseOptions(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if opts.Procs == 0 {
		opts.Procs = runtime.NumCPU()
	}

	fmt.Printf("opts:%#v\n", opts)
	fmt.Println(url)

	resp, err := request.Request("GET", url[0], "Range", "bytes=281-294")
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}
	defer resp.Body.Close()

	// getheader.Headers(os.Stdout, resp)
	dump, _ := httputil.DumpResponse(resp, false)
	fmt.Printf("response:\n%s\n", dump)

	fmt.Printf("response status code: %s\n", resp.Status)
	outs, _ := getheader.HeaderComma(os.Stdout, resp, "Accept-Ranges")
	fmt.Println(outs == "bytes")
	if err := download.DownloadFile(opts.Output, url); err != nil {
		log.Fatal(err)
	}
}
