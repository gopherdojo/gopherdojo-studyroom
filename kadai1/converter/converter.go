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

// Ext codes represents convertible file ext.
const (
	ExtJPG  = ".jpg"
	ExtJPEG = ".jpeg"
	ExtPNG  = ".png"
	ExtGIF  = ".gif"
)

// IsConvertible check ext is convertible.
func IsConvertible(ext string) bool {
	switch ext {
	case ExtJPG, ExtJPEG, ExtPNG, ExtGIF:
		return true
	default:
		return false
	}
}

// Converter converts image ext.
type Converter struct {
	FromExt        string
	ToExt          string
	TargetFilePath string
}

// Convert is main funf of Convert
func (con *Converter) Convert() error {
	var err error
	reader, err := os.Open(con.TargetFilePath)
	if err != nil {
		fmt.Printf("failed to open file: %v. error: %v", con.TargetFilePath, err)
		return err
	}
	defer reader.Close()

	// decode
	var img image.Image

	switch con.FromExt {
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
	fileName := strings.Replace(con.TargetFilePath, con.FromExt, con.ToExt, 1)
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
	switch con.ToExt {
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
