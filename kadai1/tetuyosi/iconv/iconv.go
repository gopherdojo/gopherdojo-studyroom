// package iconv implements functions to convert images to a desired format.
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

// A Convert image interface
type Conv struct {
	FilePath string
	SrcExt   string
	DestExt  string
	image    image.Image
}

// jpeg or jpg return same extension
var extMap = map[string]string{
	"jpg":  "jpg",
	"jpeg": "jpg",
	"png":  "png",
	"gif":  "gif",
}

// Return Conv object
func New(path string, srcExt string, destExt string) (*Conv, error) {
	fileExt := strings.ToLower(strings.TrimLeft(filepath.Ext(path), "."))
	srcExt = strings.ToLower(srcExt)
	destExt = strings.ToLower(destExt)
	if !isConvertible(fileExt) {
		return nil, fmt.Errorf("変換ファイル拡張子が不正です。")
	}
	if !isConvertible(srcExt) {
		return nil, fmt.Errorf("変換元フォーマット指定の誤りです。")
	}
	if !isConvertible(destExt) {
		return nil, fmt.Errorf("変換先フォーマット指定の誤りです。")
	}
	return &Conv{FilePath: path, SrcExt: srcExt, DestExt: destExt}, nil
}

// Check convertible extension
func isConvertible(ext string) bool {
	_, isExist := extMap[ext]
	return isExist
}

// Read file and convert internal image
func (c *Conv) Imaging() error {
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

// internal image convert to desired format
func (c *Conv) Convert() error {
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
