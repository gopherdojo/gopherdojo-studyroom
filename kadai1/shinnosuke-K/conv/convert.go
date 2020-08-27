package conv

import (
	"errors"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/shinnosuke-K/gopherdojo-studyroom/kadai1/shinnosuke-K/file"
)

var imgExts = []string{"gif", "png", "jpg", "jpeg"}

func Do(dirPath string, before string, after string) {

	if ok := file.ExistDir(dirPath); !ok {
		fmt.Println("not found dir")
		os.Exit(1)
	}

	if err := checkOpt(before, after); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	files := file.GetImgFiles(dirPath, before)
	for n := range files {
		if err := convert(after, files[n]); err != nil {
			log.Fatal(err)
		}
	}
}

func checkOpt(before string, after string) error {
	for n := range imgExts {
		if strings.ToLower(before) == imgExts[n] || strings.ToLower(after) == imgExts[n] {
			return nil
		}
	}
	return errors.New("imgconv: invalid image extension")
}

func convert(afterEx string, file file.File) error {
	switch afterEx {
	case "png":
		if err := convertToPNG(file); err != nil {
			return err
		}
	case "jpeg", "jpg":
		if err := convertToJPG(file); err != nil {
			return err
		}
	case "gif":
		if err := convertToGif(file); err != nil {
			return err
		}
	}
	return nil
}

func convertToPNG(f file.File) error {

	imgFile, err := file.DecodeToImg(f.Dir, f.Name)
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

func convertToJPG(f file.File) error {

	imgFile, err := file.DecodeToImg(f.Dir, f.Name)
	if err != nil {
		return err
	}

	destFileName := strings.Replace(f.Name, f.Extension, ".jpg", 1)
	destFile, err := os.Create(destFileName)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if err := jpeg.Encode(destFile, imgFile, nil); err != nil {
		return err
	}
	return nil
}

func convertToGif(f file.File) error {

	imgFile, err := file.DecodeToImg(f.Dir, f.Name)
	if err != nil {
		return err
	}

	destFileName := strings.Replace(f.Name, f.Extension, ".gif", 1)
	destFile, err := os.Create(destFileName)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if err := gif.Encode(destFile, imgFile, nil); err != nil {
		return err
	}

	return nil
}
