package kadai1

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

func (c Converter) Decode(sf io.Reader) (image.Image, error) {
	switch c.fromExt {
	case "png":
		return png.Decode(sf)
	case "jpg":
		return jpeg.Decode(sf)
	case "gif":
		return gif.Decode(sf)
	default:
		return nil, fmt.Errorf("no setting Converter.fromExt")
	}
}

func (c Converter) Encode(df io.Writer, img image.Image) error {
	switch c.toExt {
	case "png":
		return png.Encode(df, img)
	case "jpg":
		return jpeg.Encode(df, img, &jpeg.Options{Quality: 100})
	case "gif":
		return gif.Encode(df, img, &gif.Options{NumColors: 256})
	default:
		return fmt.Errorf("no setting Converter.toExt")
	}
}

func (c Converter) Create(path string) error {
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
	df, err := os.Create(fileName[:len(fileName)-len(filepath.Ext(path))] + "." + c.toExt)
	if err != nil {
		return err
	}
	defer df.Close()
	err = c.Encode(df, img)
	if err != nil {
		return nil
	}
	return nil
}

func (c Converter) Convert() {
	ext := "." + c.fromExt
	err := filepath.Walk(c.targetDir,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ext {
				c.Create(path)
			}
			return nil
		})
	if err != nil {
		fmt.Println("error :", err)
		return
	}
}

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
