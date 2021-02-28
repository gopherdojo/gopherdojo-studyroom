package image

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

//Decode,Encode周りなど、image.Imageに関するエラー
type ConvertError struct {
	SrcFormat string
	DstFormat string
	Err       error
}

//ファイルに関するエラー
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
	lastindex := strings.LastIndex(src, srcFormat)
	dst := src[:lastindex] + dstFormat

	//書き込み用ファイルを開く
	df, err := os.Create(dst)
	if err != nil {
		return &FileError{Fn: src, Err: err}
	}
	defer func() *FileError {
		if err := df.Close(); err != nil {
			return &FileError{Fn: src, Err: err}
		}
		return nil
	}()

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
