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

// 画像を変換します。
func Do(srcPath string, to Ext, rmSrc bool) error {
	// ファイルオープン
	src, openerr := os.Open(filepath.Clean(srcPath))
	if openerr != nil {
		return openerr
	}
	defer func() error {
		var srccloseerr error
		if srccloseerr := src.Close(); srccloseerr != nil {
			return srccloseerr
		}
		return srccloseerr
	}()

	// ファイルオブジェクトを画像オブジェクトに変換
	img, _, decodeerr := image.Decode(src)
	if decodeerr != nil {
		return decodeerr
	}

	// 出力ファイルを生成
	dstPath := convExt(srcPath, to)
	dst, createerr := os.Create(dstPath)
	if createerr != nil {
		return createerr
	}
	defer func() error {
		var dstcloseerr error
		if dstcloseerr = dst.Close(); dstcloseerr != nil {
			return dstcloseerr
		}
		return dstcloseerr
	}()

	// 画像ファイルを出力
	outputerr := outputImage(dst, img, to)
	if outputerr != nil {
		return outputerr
	}

	// 元ファイルを削除（オプション）
	if rmSrc {
		if removeerr := os.Remove(srcPath); removeerr != nil {
			return removeerr
		}
	}

	// 結果を標準出力
	fmt.Println("convarted " + srcPath + " to " + dstPath + "\n")

	return nil
}
