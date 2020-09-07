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

func decode(srcPath string) (image.Image, error) {
	// ファイルオープン
	src, openerr := os.Open(filepath.Clean(srcPath))
	if openerr != nil {
		return nil, openerr
	}
	var closeerr error
	defer func() {
		if closeerr = src.Close(); closeerr != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", closeerr)
		}
	}()

	// ファイルオブジェクトを画像オブジェクトに変換
	img, _, decodeerr := image.Decode(src)
	if decodeerr != nil {
		return nil, decodeerr
	}

	return img, nil
}

// 拡張子を変更
func convExt(srcPath string, to Ext) string {
	ext := filepath.Ext(srcPath)
	return srcPath[:len(srcPath)-len(ext)] + string(to)
}

func encode(dstPath string, img image.Image, to Ext) error {
	// 出力ファイルを生成
	dst, createerr := os.Create(dstPath)
	if createerr != nil {
		return createerr
	}
	defer func() {
		var closeerr error
		if closeerr = dst.Close(); closeerr != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", closeerr)
		}
	}()

	// 画像オブジェクトを指定の拡張子で出力
	outputerr := outputImage(dst, img, to)
	if outputerr != nil {
		return outputerr
	}

	return nil
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

// 画像を変換
func Do(srcPath string, to Ext, rmSrc bool) error {
	//変換前ファイルをdecode
	img, decodeerr := decode(srcPath)
	if decodeerr != nil {
		return decodeerr
	}

	//変換後ファイルのパスを生成
	dstPath := convExt(srcPath, to)

	//変換後ファイルを指定の拡張子でencode
	encodeerr := encode(dstPath, img, to)
	if encodeerr != nil {
		return encodeerr
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
