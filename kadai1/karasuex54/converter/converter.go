// Package converter is image file convert package.
package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

// Converter is a struct that has target directory, target file extention and convert file extention.
type Converter struct {
	targetDir string
	fromExt   string
	toExt     string
}

type Option func(*Converter) error

func FromExt(ext string) Option {
	return func(c *Converter) error {
		c.fromExt = ext
		return nil
	}
}

func ToExt(ext string) Option {
	return func(c *Converter) error {
		c.toExt = ext
		return nil
	}
}

// Decode reads an image from r and returns it as an image.Image.
func (c Converter) Decode(r io.Reader) (image.Image, error) {
	switch c.fromExt {
	case "png":
		return png.Decode(r)
	case "jpg", "jpeg":
		return jpeg.Decode(r)
	case "gif":
		return gif.Decode(r)
	default:
		return nil, fmt.Errorf("fail setting Converter.fromExt")
	}
}

// Encode writes the Image m to w in c.toExt format.
func (c Converter) Encode(w io.Writer, m image.Image) error {
	switch c.toExt {
	case "png":
		return png.Encode(w, m)
	case "jpg", "jpeg":
		return jpeg.Encode(w, m, &jpeg.Options{Quality: 100})
	case "gif":
		return gif.Encode(w, m, &gif.Options{NumColors: 256})
	default:
		return fmt.Errorf("fail setting Converter.toExt")
	}
}

// Convert reads an image file from path, creates a file, and writes to it in c.toExt format.
func (c Converter) Convert(path string) error {
	sf, err := os.Open(path)
	if err != nil {
		return err
	}
	defer sf.Close()
	img, err := c.Decode(sf)
	if err != nil {
		return err
	}
	fileName := filepath.Base(path)
	convertedPath := fileName[:len(fileName)-len(filepath.Ext(path))] + "." + c.toExt
	df, err := os.Create(convertedPath)
	if err != nil {
		return err
	}
	defer df.Close()
	err = c.Encode(df, img)
	if err != nil {
		if removeErr := os.Remove(convertedPath); removeErr != nil {
			return removeErr
		}
		return err
	}
	return nil
}

// Run finds files of c.fromExt form in the directories and files under c.targetDir. It passes the path of the file as an argument to c.Convert().
func (c Converter) Run() {
	ext := "." + c.fromExt
	err := filepath.Walk(c.targetDir,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ext {
				c.Convert(path)
			}
			return nil
		})
	if err != nil {
		fmt.Println("error :", err)
		return
	}
}

// NewConverter returns c.Converter.
func NewConverter(dirPath string, options ...Option) *Converter {
	c := Converter{
		targetDir: dirPath,
		fromExt:   "jpg",
		toExt:     "png",
	}
	for _, option := range options {
		option(&c)
	}
	return &c
}
