package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var  selectedDirecotry string
var  selectedFileType string
var  convertedFileType string

func returnFilePath(selectedFileType *string) ([]string,error) {
	var stringPath  []string
	err := filepath.Walk(selectedDirecotry,
		func(paths string, info fs.FileInfo, err error) error {
			if filepath.Ext(paths) == *selectedFileType {
				stringPath = append(stringPath,paths)
			}
			return nil
		})
	return  stringPath,err
}

func init()  {

	flag.StringVar(&selectedDirecotry, "s", "", "ディレクトリを指定")
	flag.StringVar(&selectedFileType, "f",".jpg", "変換前のファイルタイプを指定")
	flag.StringVar(&convertedFileType, "cf",".jpg", "変換後のファイルタイプを指定")

}

func main() {
	flag.Parse()

	paths,err := returnFilePath(&selectedFileType)
	fmt.Println(paths)

	if err != nil {
		fmt.Fprintf(os.Stderr,"ディレクトリ選択をしてください。")
		os.Exit(-1)
	}
}
