package lib

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
)

type flagStruct struct {
	selectedDirectory string
	selectedFileType  string
	convertedFileType string
	stringPath        []string
}

var flg flagStruct

func returnFilePath(selectedFileType *string) error {
	err := filepath.Walk(flg.selectedDirectory,
		func(paths string, info fs.FileInfo, err error) error {
			if filepath.Ext(paths) == *selectedFileType {
				flg.stringPath = append(flg.stringPath, paths)
			}
			return nil
		})
	return err
}
func init() {
	flag.StringVar(&flg.selectedDirectory, "s", "", "ディレクトリを指定")
	flag.StringVar(&flg.selectedFileType, "f", ".jpg", "変換前のファイルタイプを指定")
	flag.StringVar(&flg.convertedFileType, "cf", ".png", "変換後のファイルタイプを指定")

}

func Convert() {
	flag.Parse()
	err := returnFilePath(&flg.selectedFileType)
	if err == nil {
		fmt.Println(flg.stringPath)
	}
}
