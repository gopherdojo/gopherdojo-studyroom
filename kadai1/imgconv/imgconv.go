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

type ConvertImage struct {
	FromPath, ToPath string
}

func (i ConvertImage) Convert() error {
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
		err = jpeg.Encode(df, img, nil)
	case ".png":
		err = png.Encode(df, img)
	case ".gif":
		err = gif.Encode(df, img, nil)
	}
	if err != nil {
		return err
	}

	if err = os.Remove(i.FromPath); err != nil {
		return err
	}

	return nil
}

func GetConvertImages(dir, from, to string) ([]ConvertImage, error) {
	var images []ConvertImage
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+from {
			images = append(images, ConvertImage{
				FromPath: path,
				ToPath:   path[:len(path)-len(filepath.Ext(path))] + "." + to,
			})
		}
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return images, nil
}
