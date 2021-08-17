package lib

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type flagStruct struct {
	selectedDirectory string
	selectedFileType  string
	convertedFileType string
	stringPath        []string
}

var flg flagStruct

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
	if err != nil {
		return err
	}
	defer f.Close()
	//設定をデコード
	config, format, err := image.DecodeConfig(f)
	if err != nil {
		log.Fatal(err)
	}

	//フォーマット名表示
	fmt.Println("画像フォーマット：" + format)
	//サイズ表示
	fmt.Println("横幅=" + strconv.Itoa(config.Width) + ", 縦幅=" + strconv.Itoa(config.Height))
	return err
}
func Convert() {
	flag.Parse()
	err := returnFilePath(&flg.selectedFileType)
	if err != nil {
		fmt.Println("Error")
	}

	for _, v := range flg.stringPath {
		err := readImage(v)
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println(v)
	}
}
