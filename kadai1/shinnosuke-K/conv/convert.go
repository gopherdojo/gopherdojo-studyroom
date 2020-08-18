package conv

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

var imgExts = []string{"gif", "png", "jpg", "jpeg"}

func Do(dirPath string, before string, after string) {

	err := checkOpt(before, after)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	files := getImgFiles(dirPath, before)
	for n := range files {
		fmt.Println(files[n])

		switch after {
		case "png":
			if err := files[n].convertToPNG(); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func checkOpt(before string, after string) error {
	for n := range imgExts {
		if strings.ToLower(before) == imgExts[n] || strings.ToLower(after) == imgExts[n] {
			return nil
		}
	}
	return fmt.Errorf("imgconv: invaild image extension")
}

func (f *File) convertToPNG() error {

	imgFile, err := decodeImg(f.Dir, f.Name)
	if err != nil {
		return err
	}

	destFileName := strings.Replace(f.Name, f.Extension, ".png", 1)
	destFile, err := os.Create(destFileName)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if err := png.Encode(destFile, imgFile); err != nil {
		return err
	}
	return nil
}
