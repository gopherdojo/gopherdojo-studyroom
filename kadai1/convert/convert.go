package convert

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetSelectedExtensionPath(fileType string, directory string) [][]string {
	var retval [][]string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		slice := strings.Split(path, ".")
		if slice[len(slice)-1] == fileType {
			retval = append(retval, slice)
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return retval
}

func ConvertImage(fileName string, from string, to string) {
	f, err := os.Open(fileName + "." + from)
	if err != nil {
		fmt.Println("open:", err)
		return
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("decode:", err)
		return
	}

	fso, err := os.Create(fileName + "." + to)
	if err != nil {
		fmt.Println("create:", err)
		return
	}
	defer fso.Close()

	switch {
	case (from == "jpg" || from == "jpeg") || to == "png":
		jpeg.Encode(fso, img, nil)
	case from == "png" && (to == "jpg" || to == "jpeg"):
		png.Encode(fso, img)
	}

	if err := os.Remove(fileName + "." + from); err != nil {
		fmt.Println(err)
	}
}
