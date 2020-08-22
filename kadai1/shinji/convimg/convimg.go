// 画像の変換機能を提供します。
package convimg

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	_ "image/jpeg"
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

// エラー処理
func assert(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

// 画像を変換します。
func Do(srcPath string, to Ext) {
	// ファイルオープン
	src, err := os.Open(srcPath)
	assert(err, "Invalid image file path "+srcPath)
	defer src.Close()

	// ファイルオブジェクトを画像オブジェクトに変換
	img, _, err := image.Decode(src)
	assert(err, "Failed to convert source file to image.")

	// 出力ファイルを生成
	dstPath := convExt(srcPath, to)
	dst, err := os.Create(dstPath)
	assert(err, "Failed to create destination file.")
	defer dst.Close()

	// 画像ファイルを出力
	switch to {
	case PNG:
		png.Encode(dst, img)
	case JPEG:
		jpeg.Encode(dst, img, nil)
	case GIF:
		gif.Encode(dst, img, nil)
	}

	// 元ファイルを削除
	if err := os.Remove(srcPath); err != nil {
		assert(err, "Failed to delete source file")
	}

	// 結果を標準出力
	fmt.Println("convarted " + srcPath + " to " + dstPath)
}
