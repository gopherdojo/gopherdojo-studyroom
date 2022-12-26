package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 画像の拡張子を変更する
func ConvertExtensions(dirpath string, from_ex string, to_ex string) {
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if ext := strings.Replace(filepath.Ext(path), ".", "", -1); info.IsDir() || ext != from_ex {
			return nil
		}

		convert(path, to_ex)
		return nil
	})

	if err != nil {
		panic(err)
	}
}

func convert(before_path string, to_ex string) {
	before_ex := strings.Replace(filepath.Ext(before_path), ".", "", -1)
	after_path := strings.Replace(before_path, before_ex, to_ex, -1)
	f, err := os.Open(before_path)
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if err != nil {
		fmt.Println("open:", err)
		return
	}
	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("decode:", err)
		return
	}
	fso, err := os.Create(after_path)
	defer func() {
		if err := fso.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if err != nil {
		fmt.Println("create:", err)
		return
	}

	switch before_ex {
	case "jpeg", "jpg":
		err = jpeg.Encode(fso, img, &jpeg.Options{})
	case "png":
		err = png.Encode(fso, img)
	case "gif":
		err = gif.Encode(fso, img, &gif.Options{})
	default:
		fmt.Println("その変換前の形式は対応していません")
	}
	if err != nil {
		fmt.Println("encode:", err)
	}
}
