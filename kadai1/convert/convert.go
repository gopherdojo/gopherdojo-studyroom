package convert

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func GetSelectedExtensionPath(fileType string, directory string) ([]string, error) {
	var retval []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == "."+fileType {
			if f, err := os.Stat(path); !(os.IsNotExist(err) || f.IsDir()) {
				retval = append(retval, path)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return retval, nil
}

func ConvertImage(fileName string, to string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	out, err := os.Create(fileName[:len(fileName)-len(filepath.Ext(fileName))+1] + to)
	if err != nil {
		return err
	}
	defer out.Close()

	switch to {
	case "jpg", "jpeg":
		if err := jpeg.Encode(out, img, nil); err != nil {
			return err
		}
	case "png":
		if err := png.Encode(out, img); err != nil {
			return err
		}
	default:

	}

	if err := os.Remove(fileName); err != nil {
		return err
	}

	return nil
}
