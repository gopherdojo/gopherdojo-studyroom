package imgconv

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type converter struct {
	root string
	src  fileType
	dest fileType
}

func NewConverter (r string, s string, d string) converter {
	return converter{
	root: r,
	src: extToType(s),
	dest: extToType(d),
	}
}

type fileType int

const (
	jpegType = iota
	pngType
	others
)

func (ft fileType) String() string {
	switch ft {
	case jpegType:
		return "jpeg"
	case pngType:
		return "png"
	case others:
		return "other type"
	default:
		return "invalid fileType"
	}
}

func extToType (ext string) fileType {
	switch ext {
	case ".jpg", ".jpeg", "jpg", "jpeg":
		return jpegType
	case ".png", "png":
		return pngType
	default:
		return others
	}
}

func (c converter) Run() {
	err := filepath.Walk(c.root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}

		ext := filepath.Ext(path)
		ft := extToType(ext)

		if ft == jpegType {
			err = convertToPng(path)
			if err != nil {
				return err
			}
		}

		return err
	})

	if err != nil {
		fmt.Println(err)
	}
}

func convertToPng(src string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer closeFile(file)

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	dest := changeExt(src, pngType)
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer closeFile(out)

	err = png.Encode(out, img)
	if err == nil {
		fmt.Printf("converted jpeg image \"%s\" to png image \"%s\"\n",
			src, dest)
	}
	return err
}

func changeExt(path string, destExt fileType) string {
	path = strings.TrimSuffix(path, filepath.Ext(path))
	return path + "." + destExt.String()
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Printf("failed to close file: %s\n", err)
		panic(err)
	}
}
