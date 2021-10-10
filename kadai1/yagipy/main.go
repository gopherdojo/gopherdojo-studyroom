package main

import (
	"flag"
	"fmt"
	"imgconv/file_searcher"
	"imgconv/imgconv"
)

var (
	from = flag.String("from", "jpeg", "from extension")
	to = flag.String("to", "png", "after extension")
)

func init() {
	flag.Usage = func() {
		fmt.Printf(`Usage: -from FROM_FORMAT -to TO_FORMAT DIRECTORY
 		Use: convert image files,
		Default: from jpeg to png
		`)
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	dir := flag.Arg(0)

	searcher, err := file_searcher.NewFileSearcher(*from, dir)
	if err != nil {
		panic("error")
	}

	paths, err := searcher.Do()
	if err != nil {
		panic("error")
	}

	for _, path := range paths {
		convertImages(path, *to)
	}
}

func convertImages(path string, ext string) {
	conv, err := imgconv.NewImgConv(path, ext)
	if err != nil {
		panic("error")
	}
	err = conv.Do()
	if err != nil {
		panic("error")
	}
}
