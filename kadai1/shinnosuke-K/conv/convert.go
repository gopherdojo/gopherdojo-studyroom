package conv

import (
	"errors"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/shinnosuke-K/gopherdojo-studyroom/kadai1/shinnosuke-K/file"
)

func Do(dirPath string, before string, after string, delImg bool) {

	if ok := file.ExistDir(dirPath); !ok {
		fmt.Println("not found dir")
		os.Exit(1)
	}

	if err := checkOpt(before, after); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	files, err := file.GetImgFiles(dirPath, before)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for n := range files {
		if err := convert(after, files[n]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if delImg {
		if err := deleteImg(files); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

// 指定した拡張子が正しいか確認
// Check that the extension you specified is correct.
func checkOpt(before string, after string) error {
	imgExts := []string{"gif", "png", "jpg", "jpeg"}
	for n := range imgExts {
		if strings.ToLower(before) == imgExts[n] || strings.ToLower(after) == imgExts[n] {
			return nil
		}
	}
	return errors.New("image convert error: invalid image extension")
}

func convert(afterEx string, file file.File) error {
	switch afterEx {
	case "png":
		if err := convertToPNG(file); err != nil {
			return fmt.Errorf("%w:Couldn't convert to png", err)
		}
	case "jpeg", "jpg":
		if err := convertToJPG(file); err != nil {
			return fmt.Errorf("%w:Couldn't convert to jpg or jpeg", err)
		}
	case "gif":
		if err := convertToGIF(file); err != nil {
			return fmt.Errorf("%w:Couldn't convert to gif", err)
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

func convertToGIF(f file.File) error {

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

// ファイルを削除
// delete file
func deleteImg(files []file.File) error {
	for n := range files {
		path := filepath.Join(files[n].Dir, files[n].Name)
		if err := os.Remove(path); err != nil {
			return fmt.Errorf("%w: Couldn't delete %s", err, path)
		}
	}
	return nil
}
