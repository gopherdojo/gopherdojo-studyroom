package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gopherdojo/gopherdojo-studyroom/kadai1/komazz/imgconv"
)

var (
	srcExt, dstExt            string
	TooFewArgumentError       = fmt.Errorf("Too Few Arguments")
	UnsupportedExtensionError = fmt.Errorf("Unsupported extension")
)

func init() {
	flag.StringVar(&srcExt, "s", "jpg", "Optional: Extension of Source Image.")
	flag.StringVar(&dstExt, "d", "png", "Optional: Extension of Destination Image.")
	flag.Parse()
}

// 変換する候補のファイルを再帰取得する
func SrcFileList(SrcExt, srcPath string) ([]string, error) {
	var srcFileList []string
	err := filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == SrcExt {
			srcFileList = append(srcFileList, path)
		}
		return nil
	})
	return srcFileList, err
}

func exec() error {
	args := flag.Args()
	if len(args) < 1 {
		return TooFewArgumentError
	}

	c := imgconv.NewConverter(srcExt, dstExt)

	srcPath := args[0]
	dir, err := os.Stat(srcPath)
	if err != nil {
		return err
	}

	if dir.IsDir() {
		// ディレクトリ指定の場合
		fileList, err := SrcFileList(c.SrcExt, srcPath)
		if err != nil {
			return err
		}
		for _, src := range fileList {
			if err := c.Convert(src); err != nil {
				return err
			}
		}
	} else {
		// ファイル指定の場合
		if filepath.Ext(srcPath) != c.SrcExt {
			return UnsupportedExtensionError
		}
		if err := c.Convert(srcPath); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if err := exec(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
