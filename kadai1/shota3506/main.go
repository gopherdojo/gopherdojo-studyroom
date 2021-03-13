package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherdojo/gopherdojo-studyroom/kadai1/shota3506/imgcov"
)

var (
	root       = flag.String("r", ".", "root")
	srcFormat  = flag.String("sf", "jpeg", "source image format")
	destFormat = flag.String("df", "png", "destination image format")
)

func main() {
	flag.Parse()

	// initialize converter
	converter := &imgcov.Converter{
		SrcFormat:  *srcFormat,
		DestFormat: *destFormat,
	}

	err := filepath.Walk(*root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		base := strings.TrimSuffix(path, filepath.Ext(path))
		dest := base + "." + *destFormat

		ok, err := converter.Convert(path, dest)
		if err != nil {
			return err
		}
		if ok {
			fmt.Printf("convert image format: %s -> %s \n", filepath.Base(path), filepath.Base(dest))
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
