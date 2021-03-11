package conversion

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

var supportedextension = []string{"jpeg", "jpg", "png", "gif"}

//Info is a type for os.fileinfo
type Info = os.FileInfo

//ExtensionCheck is to check extension
func ExtensionCheck(extension string) error {
	for _, e := range supportedextension {
		if extension == e {
			return nil
		}
	}
	return errors.New("this extension is not supported")
}

//WalkDir is to walk directory and convert extension
func WalkDir(dir, after string) error {
	err := filepath.Walk(dir,
		func(path string, info Info, err error) error {
			if info == nil {
				return errors.New(path + " is not directory")
			}
			if os.IsNotExist(err) || info.IsDir() {
				return nil
			}

			err = convertImage(path, after)
			if err != nil {
				return err
			}

			return nil
		})

	if err != nil {
		return err
	}
	return nil
}

func convertImage(path, after string) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.New("directory is not selected")
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return errors.New("decode failed")
	}

	out, err := os.Create(strings.Replace(path, filepath.Ext(path), "."+after, 1))
	if err != nil {
		return errors.New("create failed")
	}
	defer out.Close()

	switch after {
	case "jpeg", "jpg":
		err = jpeg.Encode(out, img, nil)
		checkencode(err)
	case "png":
		err = png.Encode(out, img)
		checkencode(err)
	case "gif":
		err = gif.Encode(out, img, nil)
		checkencode(err)
	}
	return nil
}

func checkencode(err error) error {
	if err != nil {
		return errors.New("encode failed")
	}
	fmt.Println("convert successed!")
	return nil
}
