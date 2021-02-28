package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/taxintt/gopherdojo-studyroom/kadai2/taxin/converter"
)

var (
	imgFormat          string
	convertedImgFormat string
	dirPath            string
	fileFormatList     []string
)

// Pass values that is passed when we execute cli tool to variables
func init() {
	flag.StringVar(&imgFormat, "fmt", "jpg", "画像ファイルの変更前のフォーマット(jpg/jpeg/png/gif)")
	flag.StringVar(&convertedImgFormat, "outfmt", "png", "画像ファイルの変更後のフォーマット(jpg/jpeg/png/gif)")
	flag.StringVar(&dirPath, "dir", ".", "画像が配置されているディレクトリのパス")
	fileFormatList = []string{"jpg", "jpeg", "png", "gif"}
}

func main() {
	if err := validateArgs(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	eachImgDirData := converter.ImgDirData{DirPath: dirPath, ImgFormat: imgFormat, ConvertedImgFormat: convertedImgFormat}
	err := converter.WalkAndConvertImgFiles(eachImgDirData)
	if err != nil {
		log.Fatal(err)
	}
}

// validate arguments that are passed by.
//
// if you specify invalid file types or don't pass directory path that contains image files, it will raise an error.
func validateArgs() error {
	flag.Parse()
	if _, err := os.Stat(dirPath); err != nil {
		return errors.New("Error: Doesn't exists the directory that you specified")
	}
	if !validateFileFormat(imgFormat) || !validateFileFormat(convertedImgFormat) {
		return errors.New("Error: Invalid or Unsupported file format")
	}
	return nil
}

// validate types of image files
//
// if you specify invalid file types, it will return false.
func validateFileFormat(passedImgFormat string) bool {
	for _, f := range fileFormatList {
		if f == passedImgFormat {
			return true
		}
	}
	return false
}
