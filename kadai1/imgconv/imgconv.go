package imgconv

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type Converter struct {
	SrcExt string
	DstExt string
}

var (
	SupportExt = map[string]string{
		"png":  ".png",
		"jpg":  ".jpg",
		"jpeg": ".jpg",
	}
	UnkownDecodeError = fmt.Errorf("Unkown Decode Error")
	UnkownEncodeError = fmt.Errorf("Unkown Encode Error")
)

func ValidExt(ext string) bool {
	_, ok := SupportExt[ext]
	return ok
}

func NewConverter(srcExt, dstExt string) *Converter {
	return &Converter{
		SrcExt: SupportExt[srcExt],
		DstExt: SupportExt[dstExt],
	}
}

func (c *Converter) Convert(src string) error {
	// 入力ファイルを取得する
	srcfile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcfile.Close()

	// 読み出す(decode)
	img, err := c.Decode(srcfile)
	if err != nil {
		return err
	}

	// 出力ファイルを作成する
	dst := c.createDstFileName(src)
	dstfile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstfile.Close()

	// 書き出す(encode)
	err = c.Encode(dstfile, img)
	if err != nil {
		return err
	}
	return nil
}

func (c *Converter) Decode(srcfile *os.File) (image.Image, error) {
	var img image.Image
	switch c.SrcExt {
	case ".jpg":
		img, err := jpeg.Decode(srcfile)
		return img, err
	case ".png":
		img, err := png.Decode(srcfile)
		return img, err
	default:
		return img, UnkownDecodeError
	}
}

func (c *Converter) Encode(dstfile *os.File, img image.Image) error {
	switch c.DstExt {
	case ".png":
		err := png.Encode(dstfile, img)
		return err
	case ".jpg":
		err := jpeg.Encode(dstfile, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
		return err
	default:
		return UnkownEncodeError
	}
}

func (c *Converter) createDstFileName(src string) string {
	oldExt := filepath.Ext(src)
	newExt := c.DstExt
	dst := strings.Replace(src, oldExt, newExt, 1)
	return dst
}
