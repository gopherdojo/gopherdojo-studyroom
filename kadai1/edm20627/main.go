package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"path/filepath"
	"strings"

	"os"
)

var supportedFormat = []string{"png", "jpg", "jpeg", "gif"}

var from, to string

func init() {
	flag.StringVar(&from, "from", "jpg", "変更元")
	flag.StringVar(&from, "f", "jpg", "変更元(short)")
	flag.StringVar(&to, "to", "png", "変更先")
	flag.StringVar(&to, "t", "png", "変更先(short)")
}

func main() {
	flag.Parse()
	dirs := flag.Args()

	if !valid(from) || !valid(to) {
		fmt.Fprintln(os.Stderr, "supported formt is "+strings.Join(supportedFormat, ", "))
		os.Exit(1)
	}

	for _, dir := range dirs {
		err := walk(dir)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}
}

func walk(dir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if info == nil {
				return errors.New(path + " is not directory")
			}
			if info.IsDir() || filepath.Ext(path)[1:] != from {
				return nil
			}

			err = imageConvert(path)
			if err != nil {
				return err
			}

			return nil
		})

	if err != nil {
		return err
	}
	return nil
}

func imageConvert(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	out, err := os.Create(path[:len(path)-len(filepath.Ext(path))] + "." + to)
	if err != nil {
		return err
	}
	defer out.Close()

	switch to {
	case "png":
		png.Encode(out, img)
	case "jpg":
		jpeg.Encode(out, img, nil)
	case "gif":
		gif.Encode(out, img, nil)
	}

	return nil
}

func valid(extension string) bool {
	for _, v := range supportedFormat {
		if v == extension {
			return true
		}
	}
	return false
}
