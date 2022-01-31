package imgcvt

import (
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

var dir, from, to string

func init() {
	flag.StringVar(&dir, "", "", "directory to walk through")
	flag.StringVar(&from, "from", "jpg", "extension of file to convert")
	flag.StringVar(&to, "to", "png", "extension of file to convert to")
}

func Do() error {
	flag.Parse()

	err := filepath.Walk(
		flag.Arg(0),
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && filepath.Ext(path) == "."+from {
				err := readFile(path)
				if err != nil {
					return err
				}
			}
			return err
		})
	return err
}

func readFile(path string) error {
	// fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}
	fmt.Println("file was decoded")

	fso, err := os.Create(strings.TrimSuffix(path, "."+from) + "." + to)
	if err != nil {
		return err
	}
	defer fso.Close()

	png.Encode(fso, img)

	fmt.Println("File was converted")
	return err
}
