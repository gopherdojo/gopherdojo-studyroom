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

// Supported extensions.
var SupportedFormat = []string{"png", "jpg", "jpeg", "gif"}

var (
	ErrNotSpecified = errors.New("Need to specify directory or file")
	ErrMotDirectory = errors.New("Specify directory or file is not directory")
)

type ConvertImage struct {
	Filepaths    []string
	From, To     string
	DeleteOption bool
}

// Get the target files for image conversion.
func (ci *ConvertImage) Get(dirs []string) error {
	if len(dirs) == 0 {
		return ErrNotSpecified
	}

	for _, dir := range dirs {
		err := filepath.Walk(dir,
			func(path string, info os.FileInfo, err error) error {
				if info == nil {
					return ErrMotDirectory
				}
				if info.IsDir() || filepath.Ext(path)[1:] != ci.From {
					return nil
				}
				ci.Filepaths = append(ci.Filepaths, path)
				return nil
			})
		if err != nil {
			return err
		}
	}
	return nil
}

// Perform image conversion.
func (ci *ConvertImage) Convert() error {
	for _, path := range ci.Filepaths {
		err := convert(path, ci.To, ci.DeleteOption)
		if err != nil {
			return err
		}
	}
	return nil
}

// Check for supported extensions.
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

func convert(path string, to string, deleteOption bool) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	out, err := os.Create(path[:len(path)-len(filepath.Ext(path))+1] + to)
	if err != nil {
		return err
	}
	defer out.Close()

	switch to {
	case "png":
		if err := png.Encode(out, img); err != nil {
			return err
		}
	case "jpg", "jpeg":
		if err := jpeg.Encode(out, img, nil); err != nil {
			return err
		}
	case "gif":
		if err := gif.Encode(out, img, nil); err != nil {
			return err
		}
	}

	if deleteOption {
		if err := os.Remove(path); err != nil {
			return err
		}
	}

	return nil
}
