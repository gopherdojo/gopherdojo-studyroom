package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

var (
	targetDir  string
	targetExt  string
	convertExt string
	decodeFunc map[string]func(io.Reader) (image.Image, error)
	encodeFunc map[string]func(io.Writer, image.Image) error
)

func getArg() error {
	flag.StringVar(&targetExt, "te", "jpg", "target extention of image file")
	flag.StringVar(&convertExt, "ce", "png", "convert extention of image file")
	flag.Parse()
	args := flag.Args()
	switch {
	case len(args) == 0:
		return fmt.Errorf("no target directory")
	case len(args) > 1:
		return fmt.Errorf("too many args")
	}
	targetDir = args[0]
	return nil
}

func makeMapFunc() {
	decodeFunc = make(map[string]func(io.Reader) (image.Image, error))
	decodeFunc["png"] = png.Decode
	decodeFunc["jpg"] = jpeg.Decode
	decodeFunc["gif"] = gif.Decode

	encodeFunc = make(map[string]func(io.Writer, image.Image) error)
	encodeFunc["png"] = png.Encode
	encodeFunc["jpg"] = func(w io.Writer, m image.Image) error {
		return jpeg.Encode(w, m, &jpeg.Options{Quality: 100})
	}
	encodeFunc["gif"] = func(w io.Writer, m image.Image) error {
		return gif.Encode(w, m, &gif.Options{NumColors: 256})
	}
}

func convert(path string) {
	sf, err := os.Open(path)
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	defer sf.Close()

	img, err := decodeFunc[targetExt](sf)
	if err != nil {
		fmt.Println("error :", err)
		return
	}

	fileName := filepath.Base(path)
	df, err := os.Create(fileName[:len(fileName)-len(filepath.Ext(path))] + "." + convertExt)
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	defer df.Close()

	encodeFunc[convertExt](df, img)
}

func walk(dir string) {
	ext := "." + targetExt
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ext {
				convert(path)
			}
			return nil
		})
	if err != nil {
		fmt.Println("error :", err)
		return
	}
}

func MyConvert() {
	getArg()
	makeMapFunc()
	walk(targetDir)
}

func main() {
	MyConvert()
}
