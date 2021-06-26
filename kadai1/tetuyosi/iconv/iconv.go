package iconv

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

type Conv struct {
	FilePath string
	SrcExt   string
	DestExt  string
	image    image.Image
}

var extMap = map[string]string{
	"jpg":  "jpg",
	"jpeg": "jpg",
	"png":  "png",
	"gif":  "gif",
}

func New(path string, srcExt string, destExt string) *Conv {
	srcExt = strings.ToLower(srcExt)
	destExt = strings.ToLower(destExt)
	return &Conv{FilePath: path, SrcExt: srcExt, DestExt: destExt}
}

func (c *Conv) Imaging() (err error) {
	f, err := os.Open(c.FilePath)
	if err != nil {
		return fmt.Errorf("指定ファイルが開けません")
	}
	defer f.Close()

	switch extMap[c.SrcExt] {
	case "jpg":
		c.image, err = jpeg.Decode(f)
	case "png":
		c.image, err = png.Decode(f)
	case "gif":
		c.image, err = gif.Decode(f)
	default:
		return fmt.Errorf("変換前拡張子が不正です")
	}
	if err != nil {
		return fmt.Errorf("画像読み込みに失敗しました")
	}
	return nil
}

func (c *Conv) Convert() (err error) {
	name := strings.TrimSuffix(filepath.Base(c.FilePath), c.SrcExt) + c.DestExt

	f, err := os.Create(filepath.Dir(c.FilePath) + "/" + name)
	if err != nil {
		return fmt.Errorf("ファイル作成に失敗しました")
	}
	defer f.Close()
	switch extMap[c.DestExt] {
	case "jpg":
		err = jpeg.Encode(f, c.image, &jpeg.Options{})
	case "png":
		err = png.Encode(f, c.image)
	case "gif":
		err = gif.Encode(f, c.image, &gif.Options{})
	default:
		return fmt.Errorf("変換後拡張子が不正です")
	}
	if err != nil {
		return fmt.Errorf("画像変換に失敗しました")
	}
	return nil
}
