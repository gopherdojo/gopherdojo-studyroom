package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type flagStruct struct {
	selectedDirecotry string
	selectedFileType  string
	convertedFileType string
}

var flg flagStruct

func returnFilePath(selectedFileType *string) ([]string, error) {
	var stringPath []string
	err := filepath.Walk(flg.selectedDirecotry,
		func(paths string, info fs.FileInfo, err error) error {
			if filepath.Ext(paths) == *selectedFileType {
				stringPath = append(stringPath, paths)
			}
			return nil
		})
	return stringPath, err
}

func init() {

	flag.StringVar(&flg.selectedDirecotry, "s", "", "ディレクトリを指定")
	flag.StringVar(&flg.selectedFileType, "f", ".jpg", "変換前のファイルタイプを指定")
	flag.StringVar(&flg.convertedFileType, "cf", ".jpg", "変換後のファイルタイプを指定")

}

func main() {
	flag.Parse()
	paths, err := returnFilePath(&flg.selectedFileType)
	fmt.Println(paths)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ディレクトリ選択をしてください。")
		os.Exit(-1)
	}
}
