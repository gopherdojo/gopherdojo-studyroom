package conversion

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func Convert(diraName string, outDirectory string, beforeExt *string, afterExt *string) error {
	files, err := filepath.Glob(diraName + "*." + *beforeExt)
	if err != nil {
		return err
	}

	for _, file := range files {
		fileName := getFileNameWithoutExt(file)

		img, err := os.Open(file)
		if err != nil {
			return err
		}
		defer img.Close()

		config, _, err := image.Decode(img)
		if err != nil {
			return err
		}

		out, err := os.Create(outDirectory + fileName + "." + *afterExt)
		if err != nil {
			return err
		}
		defer out.Close()

		switch *afterExt {
		case "jpg":
			err := jpeg.Encode(out, config, nil)
			if err != nil {
				return err
			}
		case "png":
			err := png.Encode(out, config)
			if err != nil {
				return err
			}
		case "gif":
			err := gif.Encode(out, config, nil)
			if err != nil {
				return err
			}
		default:
		}
	}
	return err
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
