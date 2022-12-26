package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func Cmd(dirpath string, from_ex string, to_ex string) {
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if arr := strings.Split(path, "."); info.IsDir() || arr[len(arr)-1] != from_ex {
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
	arr := strings.Split(before_path, ".")
	// 拡張子以外"ドット"を使っていないと仮定
	after_path := arr[0] + "." + to_ex
	before_ex := arr[len(arr)-1]
	f, err := os.Open(before_path)
	fmt.Println(before_path, "before_path")
	defer f.Close()
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
	defer fso.Close()
	if err != nil {
		fmt.Println("create:", err)
		return
	}

	switch before_ex {
	case "jpeg", "jpg":
		jpeg.Encode(fso, img, &jpeg.Options{})
	case "png":
		png.Encode(fso, img)
	case "gif":
		gif.Encode(fso, img, &gif.Options{})
	default:
		fmt.Println("その変換前の形式は対応していません")
	}
}
