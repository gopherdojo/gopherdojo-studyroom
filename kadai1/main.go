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
	stringPath []string
}

var flg flagStruct

func returnFilePath(selectedFileType *string) (error) {
	err := filepath.Walk(flg.selectedDirecotry,
		func(paths string, info fs.FileInfo, err error) error {
			if filepath.Ext(paths) == *selectedFileType {
				flg.stringPath = append(flg.stringPath, paths)
			}
			return nil
		})
	return err
}

func init() {

	flag.StringVar(&flg.selectedDirecotry, "s", "", "ディレクトリを指定")
	flag.StringVar(&flg.selectedFileType, "f", ".jpg", "変換前のファイルタイプを指定")
	flag.StringVar(&flg.convertedFileType, "cf", ".png", "変換後のファイルタイプを指定")

}

func main() {
	flag.Parse()
	err := returnFilePath(&flg.selectedFileType)
	fmt.Println(&flg.stringPath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ディレクトリ選択をしてください。")
		os.Exit(-1)
	}
}
