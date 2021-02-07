package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/taxintt/gopherdojo-studyroom/kadai1/taxin/converter"
)

var (
	imgFormat          string
	convertedImgFormat string
	fileFormatList     []string
)

// Pass values that is passed when we execute cli tool to variables
func init() {
	flag.StringVar(&imgFormat, "f", "jpg", "画像ファイルの変更前のフォーマット(jpg/jpeg/png/gif)")
	flag.StringVar(&convertedImgFormat, "o", "png", "画像ファイルの変更後のフォーマット(jpg/jpeg/png/gif)")
	fileFormatList = []string{"jpg", "jpeg", "png", "gif"}
	flag.Parse()
}

func main() {
	dirList := flag.Args()
	if err := validateArgs(dirList); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, dir := range dirList {
		eachImgDirData := converter.ImgDirData{DirPath: dir, ImgFormat: imgFormat, ConvertedImgFormat: convertedImgFormat}
		converter.WalkAndConvertImgFiles(eachImgDirData)
	}
}

// validate arguments that are passed by.
//
// if you specify invalid file types or don't pass directory path that contains image files, it will raise an error.
func validateArgs(dirList []string) error {
	if len(dirList) == 0 {
		return errors.New("Error: Specify directory that contains image files")
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
		if f == imgFormat {
			return true
		}
	}
	return false
}
