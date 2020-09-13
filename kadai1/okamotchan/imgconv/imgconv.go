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

// 画像ファイルの拡張子
type ImageExt string

type ImageConverter struct {
	from ImageExt
	to   ImageExt
}

const (
	GIF  ImageExt = "gif"
	JPEG ImageExt = "jpeg"
	JPG  ImageExt = "jpg"
	PNG  ImageExt = "png"
)

var ValidImageExt = []ImageExt{GIF, JPEG, JPG, PNG}

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

// string型にする
func (ext ImageExt) toString() string {
	return string(ext)
}

// ImageExt型にする
func toImageExt(str string) (*ImageExt, error) {
	if str[0] == '.' {
		str = str[1:]
	}

	for _, v := range ValidImageExt {
		if v.toString() == strings.ToLower(str) {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("拡張子が正しくありません: %s", str)
}

// ImageConverter型を作成する
func newImageConverter(from, to string) (*ImageConverter, error) {
	extFrom, err := toImageExt(from)
	if err != nil {
		return nil, err
	}

	extTo, err := toImageExt(to)
	if err != nil {
		return nil, err
	}

	return &ImageConverter{from: *extFrom, to: *extTo}, nil
}

// 画像を読み込む
func (c *ImageConverter) readImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("can not open file: %s", path)
	}
	defer file.Close()

	image, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return image, nil
}

// 画像を保存する
func (c *ImageConverter) saveImage(image image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	switch c.to {
	case GIF:
		err = gif.Encode(file, image, nil)
	case JPG, JPEG:
		err = jpeg.Encode(file, image, nil)
	case PNG:
		err = png.Encode(file, image)
	default:
		err = fmt.Errorf("変換不可能な拡張子です: %s", c.to)
	}

	if err != nil {
		return err
	}

	return nil
}

// ファイル名を取得する
func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

// 画像ファイルを変換する
func (c *ImageConverter) convert(src string) error {
	dir := filepath.Dir(src)
	dst := filepath.Join(dir, fmt.Sprintf("%s.%s", getFileNameWithoutExt(src), c.to.toString()))

	if _, err := os.Stat(dst); !os.IsNotExist(err) {
		return fmt.Errorf("file already exists: %s", dst)
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
func (c *ImageConverter) convertAll(dir string) error {
	walkErr := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path)[1:] == c.from.toString() {
			if err != nil {
				if err := c.convert(path); err != nil {
					return err
				}
			}
		}
		return nil
	})

	if walkErr != nil {
		return walkErr
	}
	return nil
}
