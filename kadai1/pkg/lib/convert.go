package lib

import (
	"flag"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/fs"
	"os"
	"path/filepath"
)

type flagStruct struct {
	selectedDirectory string
	selectedFileType  string
	convertedFileType string
	stringPath        []string
}

var flg flagStruct

func assert(err error) error {
	if err != nil {
		fmt.Println("error")
		return err
	}
	return err
}
func init() {
	flag.StringVar(&flg.selectedDirectory, "s", "", "ディレクトリを指定")
	flag.StringVar(&flg.selectedFileType, "f", ".jpg", "変換前のファイルタイプを指定")
	flag.StringVar(&flg.convertedFileType, "cf", ".png", "変換後のファイルタイプを指定")

}

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

func readImage(fn string) error {
	f, err := os.Open(fn)
	err = assert(err)
	defer f.Close()
	return err
}
func Convert() {
	flag.Parse()
	err := returnFilePath(&flg.selectedFileType)
	err = assert(err)

	for _, v := range flg.stringPath {
		err := readImage(v)
		err = assert(err)
		fmt.Println(v)
	}

}
