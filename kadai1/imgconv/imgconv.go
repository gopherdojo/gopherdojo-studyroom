package imgconv

import (
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type TargetImage struct {
	fromPath string
	toPath   string
}

func (i TargetImage) Convert() error {
	sf, err := os.Open(i.fromPath)
	if err != nil {
		return err
	}
	defer sf.Close()

	img, _, err := image.Decode(sf)
	if err != nil {
		return err
	}

	df, _ := os.Create(i.toPath)
	defer df.Close()

	switch filepath.Ext(i.toPath) {
	case ",png":
		err = png.Encode(df, img)
	}
	if err != nil {
		return err
	}

	err = os.Remove(i.fromPath)
	if err != nil {
		return err
	}

	return nil
}

func FileWalk(dir string, from string, to string) ([]TargetImage, error) {
	var images []TargetImage
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+from {
			images = append(images, TargetImage{
				fromPath: path,
				toPath:   path[:len(path)-len(filepath.Ext(path))] + "." + to,
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
