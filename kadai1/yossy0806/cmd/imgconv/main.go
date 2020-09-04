package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	imgconv "github.com/yossy0806/gopherdojo-studyroom/kadai1/yossy0806/imgconv"
)

var (
	dir string
	se  string
	de  string
)

func init() {
	flag.StringVar(&dir, "dir", "", "変換対象のディレクトリの指定")
	flag.StringVar(&se, "se", "jpg", "変更前の画像ファイルの指定(jpg|jpeg|png|gif)")
	flag.StringVar(&de, "de", "png", "変更後の画像ファイルの指定(jpg|jpeg|png|gif)")
	flag.Parse()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

func run() error {
	if dir == "" {
		return errors.New("please specify a directory")
	}

	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return fmt.Errorf("the specified directory was not found: dir=%s", dir)
	}

	if info.IsDir() {
		err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Errorf("access to path failed: path=%q: %v", path, err)
				return err
			}

			if info.IsDir() {
				return nil
			}

			converter := imgconv.NewConverter(path, se, de)
			if err := converter.Validate(); err != nil {
				return err
			}

			// 拡張子がflagの引数で指定できるもの、defaultの拡張子でなければskip
			ext := strings.ToLower(filepath.Ext(path))
			if ext != "."+se {
				return nil
			}

			img, err := converter.Decode()
			if err != nil {
				return err
			}

			if err := converter.Encode(img); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	fmt.Println("the image conversion was successful.")
	return nil
}
