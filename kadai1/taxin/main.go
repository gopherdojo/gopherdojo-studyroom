package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/taxintt/gopherdojo-studyroom/kadai1/taxin/converter"
)

var (
	imgFormat          string
	convertedImgFormat string
	dirPath            string
	fileFormatList     []string
)

// Pass values that is passed when we execute cli tool to variables
func init() {
	flag.StringVar(&imgFormat, "f", "jpg", "画像ファイルの変更前のフォーマット(jpg/jpeg/png/gif)")
	flag.StringVar(&convertedImgFormat, "o", "png", "画像ファイルの変更後のフォーマット(jpg/jpeg/png/gif)")
	flag.StringVar(&dirPath, "d", ".", "画像が配置されているディレクトリのパス")
	fileFormatList = []string{"jpg", "jpeg", "png", "gif"}
	flag.Parse()
}

func main() {
	if err := validateArgs(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	eachImgDirData := converter.ImgDirData{DirPath: dirPath, ImgFormat: imgFormat, ConvertedImgFormat: convertedImgFormat}
	converter.WalkAndConvertImgFiles(eachImgDirData)
}

// validate arguments that are passed by.
//
// if you specify invalid file types or don't pass directory path that contains image files, it will raise an error.
func validateArgs() error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	if existsDir(files) {
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
		if f == imgFormat {
			return true
		}
	}
	return false
}

func existsDir(dirs []os.FileInfo) bool {
	for _, f := range dirs {
		if f.Name() == dirPath {
			return true
		}
	}
	return false
}
