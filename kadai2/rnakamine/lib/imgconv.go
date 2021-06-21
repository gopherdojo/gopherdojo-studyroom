package imgconv

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// ConvertImage is information about the path of the image to be converted.
type ConvertImage struct {
	FromPath, ToPath string
}

// Convert converts an image to a specified format
func (i *ConvertImage) Convert(deleteOption bool) error {
	sf, err := os.Open(i.FromPath)
	if err != nil {
		return err
	}
	defer sf.Close()

	img, _, err := image.Decode(sf)
	if err != nil {
		return err
	}

	df, err := os.Create(i.ToPath)
	if err != nil {
		return err
	}
	defer df.Close()

	switch strings.ToLower(filepath.Ext(i.ToPath)) {
	case ".jpeg", ".jpg":
		err = jpeg.Encode(df, img, &jpeg.Options{Quality: 100})
	case ".png":
		err = png.Encode(df, img)
	case ".gif":
		err = gif.Encode(df, img, nil)
	}
	if err != nil {
		return err
	}

	if deleteOption {
		if err = os.Remove(i.FromPath); err != nil {
			return err
		}
	}

	return nil
}

// GetConvertImages retrieves the images to be converted from the specified directory
func GetConvertImages(dir, from, to string) ([]ConvertImage, error) {
	var images []ConvertImage
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == "."+from && !info.IsDir() {
			images = append(images, ConvertImage{
				FromPath: path,
				ToPath:   path[:len(path)-len(filepath.Ext(path))] + "." + to,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return images, nil
}
