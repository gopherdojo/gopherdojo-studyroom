package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	iconv "github.com/tetuyosi/gopherdojo-studyroom/kadai1/tetuyosi/iconv"
)

// 変換対象ディレクトリ名
var dir string

// 変換対象拡張子(jpg,jpeg,png,gif)
var srcExt string

// 変換後拡張子(jpg,jpeg,png,gif)
var destExt string

func init() {
	flag.StringVar(&srcExt, "srcExt", "jpg", "変換前画像フォーマット")
	flag.StringVar(&destExt, "destExt", "png", "変換後画像フォーマット")
}

func main() {
	flag.Parse()

	dir = flag.Arg(0)
	// 引数のチェック
	if dir == "" {
		fmt.Printf("変換対象ディレクトリ名を入れてください。")
		os.Exit(1)
	}
	_, err := os.Stat(dir)
	// フォルダチェック
	if os.IsNotExist(err) {
		fmt.Printf("指定ディレクトリは存在しません。%s", dir)
		os.Exit(1)
	}

	// 画像変換
	err = filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		c, err := iconv.New(path, srcExt, destExt)
		if err != nil {
			return nil
		}
		err = c.Imaging()
		if err != nil {
			return err
		}
		err = c.Convert()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	os.Exit(0)
}
