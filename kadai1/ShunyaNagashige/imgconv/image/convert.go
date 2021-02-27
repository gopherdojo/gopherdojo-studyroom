package image

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

type ConvertError struct {
	SrcFormat string
	DstFormat string
	Err       error
}

type FileError struct {
	Fn  string
	Err error
}

func (err *ConvertError) Error() string {
	return err.Err.Error()
}

func (err *FileError) Error() string {
	return err.Err.Error()
}

func Convert(src, srcFormat, dstFormat string) error {

	if srcFormat == "jpeg" {
		srcFormat = "jpg"
	}

	if dstFormat == "jpeg" {
		dstFormat = "jpg"
	}

	//読み込み用にファイルを開く
	sf, err := os.Open(src)
	if err != nil {
		return &FileError{Fn: src, Err: err}
	}
	defer sf.Close()

	img, _, err := image.Decode(sf)
	if err != nil {
		return &ConvertError{DstFormat: dstFormat, SrcFormat: srcFormat, Err: err}
	}

	//変換後のファイルのパスを作成
	dst := strings.Replace(src, "."+srcFormat, "."+dstFormat, 1)

	//書き込み用ファイルを開く
	df, err := os.Create(dst)
	if err != nil {
		return &FileError{Fn: src, Err: err}
	}
	defer df.Close()

	//目的の形式にエンコード
	switch dstFormat {
	case "png":
		if err := png.Encode(df, img); err != nil {
			return &ConvertError{DstFormat: dstFormat, SrcFormat: srcFormat, Err: err}
		}
	case "jpg", "jpeg":
		if err := jpeg.Encode(df, img, nil); err != nil {
			return &ConvertError{DstFormat: dstFormat, SrcFormat: srcFormat, Err: err}
		}
	case "gif":
		if err := gif.Encode(df, img, nil); err != nil {
			return &ConvertError{DstFormat: dstFormat, SrcFormat: srcFormat, Err: err}
		}
	}

	return nil
}
