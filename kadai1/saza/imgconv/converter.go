package imgconv

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
)

type Converter struct {
	Root string
}

func (c Converter) Run() {
	err := filepath.Walk(c.Root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}

		fmt.Println(path)

		ext := filepath.Ext(path)

		if ext == ".jpg" || ext == ".jpeg" {
			err = convertToPng(path, path + ".png")
			if err != nil {
				fmt.Println("failed to load jpeg")
			}
		}

		return err
	})

	if err != nil {
		fmt.Println(err)
	}
}

func convertToPng(src string, dest string) error {
	file, err := os.Open(src)
	fmt.Println(err)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	out, err := os.Create(dest)
	defer out.Close()

	err = png.Encode(out, img)
	return err
}
