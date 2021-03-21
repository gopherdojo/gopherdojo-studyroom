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

const (
	JPEG string = "jpeg"
	JPG  string = "jpg"
	PNG  string = "png"
	GIF  string = "gif"
)

type ConvertStruct struct {
	filepaths     []string
	Before, After string
}

//ExtensionCheck is to check extension
func ExtensionCheck(extension string) error {
	switch extension {
	case JPEG, JPG, PNG, GIF:
		return nil
	default:
		return errors.New("this extension is not supported")
	}
}

func (cs *ConvertStruct) WalkDirs(dirs []string) error {
	for _, dir := range dirs {
		err := filepath.Walk(dir,
			func(path string, info os.FileInfo, err error) error {
				if info == nil {
					return errors.New(path + " is not directory")
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
			return errors.New("convert failed")
		}
	}
	return nil
}

func convert(path, after string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	out, err := os.Create(strings.Replace(path, filepath.Ext(path), "."+after, 1))
	if err != nil {
		return err
	}
	defer out.Close()

	switch after {
	case JPEG, JPG:
		if err := jpeg.Encode(out, img, nil); err != nil {
			return err
		}
		fmt.Println("convert successed")
		return nil
	case PNG:
		if err := png.Encode(out, img); err != nil {
			return err
		}
		fmt.Println("convert successed")
		return nil
	case GIF:
		if err := gif.Encode(out, img, nil); err != nil {
			return err
		}
		fmt.Println("convert successed")
		return nil
	default:
		return err
	}
}
