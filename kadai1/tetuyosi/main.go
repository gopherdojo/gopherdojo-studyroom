package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var directory string
var fromFormat string
var toFormat string

func init() {
	flag.StringVar(&fromFormat, "from", "jpg", "変換前画像フォーマット")
	flag.StringVar(&toFormat, "to", "png", "変換後画像フォーマット")
}

func main() {
	flag.Parse()

	directory = flag.Arg(0)
	// 引数のチェック
	if directory == "" {
		fmt.Fprintf(os.Stderr, "%s", "変換対象ディレクトリ名を入れてください。")
		os.Exit(1)
	}
	_, err := os.Stat(directory)
	// フォルダチェック
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s%s", "指定ディレクトリは存在しません。", directory)
		os.Exit(1)
	}

	// ファイル名表示
	filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), fromFormat) {
			return nil
		}
		fmt.Printf("name: %+v \n", info.Name())
		target := strings.TrimSuffix(info.Name(), fromFormat) + toFormat
		fmt.Printf("target: %+v \n", target)

		rf, errr := os.Open(path)
		if errr != nil {
			return errr
		}
		defer rf.Close()

		img, errr := jpeg.Decode(rf)
		if errr != nil {
			return errr
		}

		f, errr := os.Create(target)
		if errr != nil {
			return errr
		}
		defer f.Close()

		errr = png.Encode(f, img)
		if errr != nil {
			return errr
		}
		fmt.Printf("the image : %s conversion was successful.", target)
		return nil
	})
	os.Exit(0)
}
