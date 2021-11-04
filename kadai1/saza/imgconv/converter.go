package imgconv

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type Converter struct {
	Root string
}

type fileType int

const (
	jpegType = iota
	pngType
	others
)

func String (ft fileType) string {
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
	case ".jpg", ".jpeg":
		return jpegType
	case ".png":
		return pngType
	default:
		return others
	}
}

func (c Converter) Run() {
	err := filepath.Walk(c.Root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}

		ext := filepath.Ext(path)
		ft := extToType(ext)

		if ft == jpegType {
			err = convertToPng(path, path + ".png")
			if err != nil {
				fmt.Println("failed to load jpeg")
			}
		}

		return err
	})

	if err != nil {
		fmt.Println(err)
	}
}

func convertToPng(src string, dest string) error {
	file, err := os.Open(src)
	fmt.Println(err)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	out, err := os.Create(dest)
	defer out.Close()

	err = png.Encode(out, img)
	return err
}
