package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

var supportedFormat = []string{"png", "jpg", "jpeg", "gif"}

type convertImage struct {
	filepaths []string
	from, to  string
}

func (ci *convertImage) get(dirs []string) error {
	for _, dir := range dirs {
		err := filepath.Walk(dir,
			func(path string, info os.FileInfo, err error) error {
				if info == nil {
					return errors.New(path + " is not directory")
				}
				if info.IsDir() || filepath.Ext(path)[1:] != ci.from {
					return nil
				}
				ci.filepaths = append(ci.filepaths, path)
				return nil
			})
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (ci *convertImage) convert() error {
	for _, path := range ci.filepaths {
		err := convert(path, ci.to)
		if err != nil {
			return err
		}
	}
	return nil
}

func convert(path string, to string) error {
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
	case "jpg", "jpeg":
		jpeg.Encode(out, img, nil)
	case "gif":
		gif.Encode(out, img, nil)
	}
	return nil
}

var ci = convertImage{}

func init() {
	flag.StringVar(&ci.from, "from", "jpg", "変更元")
	flag.StringVar(&ci.from, "f", "jpg", "変更元(short)")
	flag.StringVar(&ci.to, "to", "png", "変更先")
	flag.StringVar(&ci.to, "t", "png", "変更先(short)")
}

func main() {
	flag.Parse()
	dirs := flag.Args()

	if !valid(ci.from) || !valid(ci.to) {
		fmt.Fprintln(os.Stderr, "supported formt is "+strings.Join(supportedFormat, ", "))
		os.Exit(1)
	}

	err := ci.get(dirs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	err = ci.convert()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func valid(extension string) bool {
	for _, v := range supportedFormat {
		if v == extension {
			return true
		}
	}
	return false
}
