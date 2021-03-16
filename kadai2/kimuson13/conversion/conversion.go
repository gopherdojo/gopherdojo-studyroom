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

type ConvertStruct struct {
	filepaths     []string
	Before, After string
}

//ExtensionCheck is to check extension
func ExtensionCheck(extension string) error {
	for _, e := range supportedextension {
		if extension == e {
			return nil
		}
	}
	return errors.New("this extension is not supported")
}

func (cs *ConvertStruct) WalkDirs(dirs []string) error {
	for _, dir := range dirs {
		err := filepath.Walk(dir,
			func(path string, info os.FileInfo, err error) error {
				if info == nil {
					return errors.New(path + " is not direcotory")
				}
				if os.IsNotExist(err) || info.IsDir() {
					return nil
				}
				cs.filepaths = append(cs.filepaths, path)
				return nil
			})
		if err != nil {
			return err
		}
	}
	return nil
}

func (cs *ConvertStruct) Convert() error {
	for _, path := range cs.filepaths {
		err := convert(path, cs.After)
		if err != nil {
			return err
		}
	}
	return nil
}

func convert(path, after string) error {
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
		if err := jpeg.Encode(out, img, nil); err != nil {
			return errors.New("encode failed")
		}
		fmt.Println("convert successed")
		return nil
	case "png":
		if err := png.Encode(out, img); err != nil {
			return errors.New("encode failed")
		}
		fmt.Println("convert successed")
		return nil
	case "gif":
		if err := gif.Encode(out, img, nil); err != nil {
			return errors.New("encode failed")
		}
		fmt.Println("convert successed")
		return nil
	}
	return nil
}
