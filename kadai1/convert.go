package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dai65527/gopherdojo-studyroom/kadai1/imageConverter"
)

var input = flag.String("i", "jpg", "input type (jpg or png, default: jpg)")
var output = flag.String("o", "png", "jpg or png (jpg or png, default: png)")

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "error: invalid argument")
		os.Exit(1)
	} else if _, err := os.Stat(flag.Args()[0]); err != nil {
		fmt.Fprintln(os.Stderr, "error: "+flag.Args()[0]+": no such file or directory")
		os.Exit(1)
	}
	convert(flag.Args()[0], isReverse())
}

func isReverse() bool {
	if *input == "png" && (*output == "jpg" || *output == "jpeg") {
		return true
	} else if (*input == "jpg" || *input == "jpeg") && *output == "png" {
		return false
	} else {
		fmt.Fprintln(os.Stderr, "error: invalid option")
		os.Exit(1)
	}
	return false
}

func convert(dirpath string, isRev bool) {
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if isRev {
				imageConverter.Png2jpg(path)
			} else {
				imageConverter.Jpg2png(path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
