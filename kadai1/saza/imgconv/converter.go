package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type converter struct {
	root string
	src  fileType
	dest fileType
}

// Converter provides a function that convert image file format.
type Converter interface {
	Run()
}

// NewConverter creates a new Converter object,
// specifying root directory, src image file format, and dest image file format.
func NewConverter(root string, src string, dest string) Converter {
	return converter{
		root: root,
		src:  extToType(src),
		dest: extToType(dest),
	}
}

type fileType int

const (
	jpegType = iota
	pngType
	gifType
	others
)

func (ft fileType) String() string {
	switch ft {
	case jpegType:
		return "jpg"
	case pngType:
		return "png"
	case gifType:
		return "gif"
	case others:
		return "other type"
	default:
		return "invalid fileType"
	}
}

func extToType(ext string) fileType {
	switch ext {
	case ".jpg", ".jpeg", "jpg", "jpeg":
		return jpegType
	case ".png", "png":
		return pngType
	case ".gif", "gif":
		return gifType
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

		if ft == c.src {
			err = convert(path, c.src, c.dest)
			if err != nil {
				return err
			}
		}

		return err
	})

	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
}

func convert(src string, srcType fileType, destType fileType) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer closeFile(file)

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	dest := changeExt(src, destType)
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer closeFile(out)

	err = encode(out, img, destType)

	if err == nil {
		fmt.Printf("converted %s image \"%s\" to %s image \"%s\"\n",
			srcType, src, destType, dest)
	}

	return err
}

func changeExt(path string, destExt fileType) string {
	path = strings.TrimSuffix(path, filepath.Ext(path))
	return path + "." + destExt.String()
}

func encode(out io.Writer, img image.Image, dest fileType) error {
	var err error

	switch dest {
	case jpegType:
		err = jpeg.Encode(out, img, &jpeg.Options{})
	case pngType:
		err = png.Encode(out, img)
	case gifType:
		err = gif.Encode(out, img, &gif.Options{})
	default:
		panic("unknown output file type")
	}

	return err
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Printf("failed to close file: %s\n", err)
		panic(err)
	}
}
