package imgcvt

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

type ConvertParams struct {
	Dir  string
	From string
	To   string
}

func Convert(p ConvertParams) error {
	flag.Parse()

	var err error
	err = validateFlag(p.From, p.To)
	if err != nil {
		return err
	}

	err = filepath.Walk(
		flag.Arg(0),
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && filepath.Ext(path) == "."+p.From {
				err := convertImage(p.Dir, p.From, p.To)
				if err != nil {
					return err
				}
			}
			return err
		})

	if err != nil {
		return err
	}

	fmt.Println("Finished converting images")

	return nil
}

func validateFlag(from, to string) error {
	if from != "png" && from != "jpeg" && from != "jpg" && from != "gif" {
		return errors.New("`from` flag is invalid")
	}

	if to != "png" && to != "jpeg" && to != "jpg" && to != "gif" {
		return errors.New("`to` flag is invalid")
	}

	return nil
}

func convertImage(from, to, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var img image.Image

	switch from {
	case "jpg", "jpeg":
		img, err = jpeg.Decode(file)
	case "png":
		img, err = png.Decode(file)
	case "gif":
		img, err = gif.Decode(file)
	default:
		err = errors.New("invalid extension")
	}

	if err != nil {
		return err
	}

	fso, err := os.Create(strings.TrimSuffix(path, "."+from) + "." + to)
	if err != nil {
		return err
	}
	defer fso.Close()

	switch to {
	case "jpg", "jpeg":
		err = jpeg.Encode(fso, img, &jpeg.Options{})
	case "png":
		err = png.Encode(fso, img)
	case "gif":
		err = gif.Encode(fso, img, nil)
	default:
		err = errors.New("invalid extension")
	}

	if err != nil {
		return err
	}

	return nil
}
