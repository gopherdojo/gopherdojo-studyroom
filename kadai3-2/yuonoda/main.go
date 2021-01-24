package main

import (
	"flag"
	splitDownload "github.com/yuonoda/gopherdojo-studyroom/kadai3-2/yuonoda/lib"
)

var url = flag.String("url", "", "URL to download")
var splitCount = flag.Int("c", 1, "how many times you split content")

func main() {
	flag.Parse()
	splitDownload.Run(*url, *splitCount)
}
