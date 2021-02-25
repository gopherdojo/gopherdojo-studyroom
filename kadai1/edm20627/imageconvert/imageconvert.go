package imageconvert

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

var SupportedFormat = []string{"png", "jpg", "jpeg", "gif"}

type ConvertImage struct {
	filepaths []string
	From, To  string
}

func (ci *ConvertImage) Get(dirs []string) error {
	for _, dir := range dirs {
		err := filepath.Walk(dir,
			func(path string, info os.FileInfo, err error) error {
				if info == nil {
					return errors.New(path + " is not directory")
				}
				if info.IsDir() || filepath.Ext(path)[1:] != ci.From {
					return nil
				}
				ci.filepaths = append(ci.filepaths, path)
				return nil
			})
		if err != nil {
			return err
		}
	}
	return nil
}

func (ci *ConvertImage) Convert() error {
	for _, path := range ci.filepaths {
		err := convert(path, ci.To)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ci *ConvertImage) Valid() bool {
	for _, v := range SupportedFormat {
		if v == ci.From {
			for _, v := range SupportedFormat {
				if v == ci.To {
					return true
				}
			}
		}
	}
	return false
}

func convert(path string, To string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	out, err := os.Create(path[:len(path)-len(filepath.Ext(path))+1] + To)
	if err != nil {
		return err
	}
	defer out.Close()

	switch To {
	case "png":
		png.Encode(out, img)
	case "jpg", "jpeg":
		jpeg.Encode(out, img, nil)
	case "gif":
		gif.Encode(out, img, nil)
	}
	return nil
}
