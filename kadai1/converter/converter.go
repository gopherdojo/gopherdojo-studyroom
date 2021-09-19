package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// Ext codes represents convertible file ext
const (
	ExtJPG  = ".jpg"
	ExtJPEG = ".jpeg"
	ExtPNG  = ".png"
	ExtGIF  = ".gif"
)

// IsConvertible check that specified ext is convertible
func IsConvertible(ext string) bool {
	switch ext {
	case ExtJPG, ExtJPEG, ExtPNG, ExtGIF:
		return true
	default:
		return false
	}
}

// Converter is ...
type Converter struct {
	FromExt        string
	ToExt          string
	TargetFilePath string
}

// Convert is a main func of image convert
func (c *Converter) Convert() error {
	var err error

	reader, err := os.Open(c.TargetFilePath)
	if err != nil {
		fmt.Printf("failed to open file: %v. error: %v", c.TargetFilePath, err)
		return err
	}
	defer reader.Close()

	// decode
	var img image.Image

	switch c.FromExt {
	case ExtJPG, ExtJPEG:
		img, err = jpeg.Decode(reader)
	case ExtPNG:
		img, err = png.Decode(reader)
	case ExtGIF:
		img, err = gif.Decode(reader)
	}

	if err != nil {
		fmt.Printf("failed to decode: %v\n", err)
		return err
	}

	// create dist file
	fileName := strings.Replace(c.TargetFilePath, c.FromExt, c.ToExt, 1)
	dist, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func() {
		if err := dist.Close(); err != nil {
			fmt.Printf("failed to close file: %v\n", err)
		}
	}()

	// encode
	switch c.ToExt {
	case ExtJPG, ExtJPEG:
		err = jpeg.Encode(dist, img, &jpeg.Options{})
	case ExtPNG:
		err = png.Encode(dist, img)
	case ExtGIF:
		err = gif.Encode(dist, img, &gif.Options{})
	}

	if err != nil {
		return err
	}
	return nil
}
