// 画像の変換機能を提供します。
package convimg

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type Ext string

const (
	PNG  Ext = ".png"
	JPEG Ext = ".jpg"
	GIF  Ext = ".gif"
)

// 拡張子を変更します。
func convExt(srcPath string, to Ext) string {
	ext := filepath.Ext(srcPath)
	return srcPath[:len(srcPath)-len(ext)] + string(to)
}

// 画像ファイルを出力
func outputImage(dst *os.File, img image.Image, to Ext) error {
	switch to {
	case PNG:
		err := png.Encode(dst, img)
		return err
	case JPEG:
		err := jpeg.Encode(dst, img, nil)
		return err
	case GIF:
		err := gif.Encode(dst, img, nil)
		return err
	default:
		return nil
	}
}

// エラー処理
func assert(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

// 画像を変換します。
func Do(srcPath string, to Ext, rmSrc bool) {
	// ファイルオープン
	src, err := os.Open(filepath.Clean(srcPath))

	assert(err, "Invalid image file path "+srcPath)

	defer func() {
		if err := src.Close(); err != nil {
			assert(err, "Failed to close destinatiln file")
		}
	}()

	// ファイルオブジェクトを画像オブジェクトに変換
	img, _, err := image.Decode(src)
	assert(err, "Failed to convert source file to image.")

	// 出力ファイルを生成
	dstPath := convExt(srcPath, to)
	dst, err := os.Create(dstPath)
	assert(err, "Failed to create destination file.")

	defer func() {
		if err := dst.Close(); err != nil {
			assert(err, "Failed to close destinatiln file")
		}
	}()

	// 画像ファイルを出力
	err = outputImage(dst, img, to)
	assert(err, "Failed to output image file.\n")

	// 元ファイルを削除（オプション）
	if rmSrc {
		if err := os.Remove(srcPath); err != nil {
			assert(err, "Failed to delete source file")
		}
	}

	// 結果を標準出力
	fmt.Println("convarted " + srcPath + " to " + dstPath)
}
