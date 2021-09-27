package kadai1

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

type Converter struct {
	targetDir string
	fromExt   string
	toExt     string
}

func (c Converter) Init() error {
	flag.StringVar(&c.fromExt, "fe", "jpg", "target extention of image file")
	flag.StringVar(&c.toExt, "te", "png", "convert extention of image file")
	flag.Parse()
	args := flag.Args()
	switch {
	case len(args) == 0:
		return fmt.Errorf("no target directory")
	case len(args) > 1:
		return fmt.Errorf("too many args")
	}
	c.targetDir = args[0]
	return nil
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
		return png.Decode(sf)
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
		return png.Encode(df, img)
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
