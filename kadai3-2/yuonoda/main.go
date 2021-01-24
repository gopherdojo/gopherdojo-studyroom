package main

import (
	splitDownload "github.com/yuonoda/gopherdojo-studyroom/kadai3-2/yuonoda/lib"
)

func main() {
	url := "https://dumps.wikimedia.org/jawiki/20210101/jawiki-20210101-pages-articles-multistream-index.txt.bz2"
	splitDownload.Run(url, 1000000)
}
