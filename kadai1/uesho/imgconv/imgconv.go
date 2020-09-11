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

	"golang.org/x/image/bmp"
)

type ImageExt string // 画像ファイルの拡張子

type imageConverter struct {
	from ImageExt
	to   ImageExt
}

const (
	GIF  ImageExt = "gif"
	JPEG ImageExt = "jpeg"
	JPG  ImageExt = "jpg"
	PNG  ImageExt = "png"
	BMP  ImageExt = "bmp"
)

var ValidImageExts = []ImageExt{GIF, JPEG, JPG, PNG, BMP}

// 画像を変換する
func Do(args map[string]string) error {
	converter, err := newImageConverter(args["from"], args["to"])
	if err != nil {
		return err
	}

	path := args["dir"]

	err = converter.convertAll(path)
	if err != nil {
		return err
	}

	return nil
}

// String型にする
func (ext ImageExt) toString() string {
	return string(ext)
}

// ImageExt型にする
func toImageExt(str string) (*ImageExt, error) {
	if str[0] == '.' {
		str = str[1:]
	}

	for _, v := range ValidImageExts {
		if v.toString() == strings.ToLower(str) {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("拡張子が正しくありません: %s", str)
}

// imageConverter型を作成する
func newImageConverter(from, to string) (*imageConverter, error) {
	extFrom, err := toImageExt(from)
	if err != nil {
		return nil, err
	}

	extTo, err := toImageExt(to)
	if err != nil {
		return nil, err
	}

	return &imageConverter{from: *extFrom, to: *extTo}, nil
}

// 画像を読み込む
func (c *imageConverter) readImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("ファイルを開くことができませんでした: %s", path)
	}
	defer file.Close()

	image, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return image, nil
}

// 画像を保存する
func (c *imageConverter) saveImage(image image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	switch c.to {
	case GIF:
		err = gif.Encode(file, image, nil)
	case JPEG, JPG:
		err = jpeg.Encode(file, image, nil)
	case PNG:
		err = png.Encode(file, image)
	case BMP:
		err = bmp.Encode(file, image)
	default:
		err = fmt.Errorf("変換不可能な拡張子です: %s", c.to)
	}

	if err != nil {
		return err
	}

	return nil
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

// 画像ファイルを変換する
func (c *imageConverter) convert(src string) error {
	dir := filepath.Dir(src)
	dst := filepath.Join(dir, fmt.Sprintf("%s.%s", getFileNameWithoutExt(src), c.to.toString()))

	if _, err := os.Stat(dst); !os.IsNotExist(err) {
		return fmt.Errorf("ファイルはすでに存在しています: %s", dst)
	}

	img, err := c.readImage(src)
	if err != nil {
		return err
	}

	err = c.saveImage(img, dst)
	if err != nil {
		return err
	}

	return nil
}

// ディレクトリ内の指定された画像ファイルを全て変換する
func (c *imageConverter) convertAll(dir string) error {
	walk_err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path)[1:] == c.from.toString() {
			if err := c.convert(path); err != nil {
				return err
			}
		}

		return nil
	})

	if walk_err != nil {
		return walk_err
	}

	return nil
}
