package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	iconv "github.com/tetuyosi/gopherdojo-studyroom/kadai1/tetuyosi/iconv"
)

var dir string
var srcExt string
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
		fmt.Fprintf(os.Stderr, "%s", "変換対象ディレクトリ名を入れてください。")
		os.Exit(1)
	}
	_, err := os.Stat(dir)
	// フォルダチェック
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s%s", "指定ディレクトリは存在しません。", dir)
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
		if !strings.HasSuffix(info.Name(), srcExt) {
			return nil
		}
		c := iconv.New(path, srcExt, destExt)
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
