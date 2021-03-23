//画像変換
package conversion

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func Convert(srcPath string, output_fmt string) error {
	//画像のフォーマットによって変換方法を変える
	//ファイル読み込み
	// ファイルオープン
	file, err := os.Open(srcPath)
	if err != nil {
		println("Error when opening a file")
		return err
	}

	defer file.Close()

	// ファイルオブジェクトを画像オブジェクトに変換
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	//出力ファイル生成
	output_path := srcPath[:len(srcPath)-len(filepath.Ext(srcPath))] + "." + output_fmt
	out, err := os.Create(output_path)
	if err != nil {
		return err
	}
	defer out.Close()
	//ファイル出力(変換)

	switch filepath.Ext(output_path) {
	case ".png":
		png.Encode(out, img)
	case ".jpeg":
		opts := &jpeg.Options{Quality: 100}
		jpeg.Encode(out, img, opts)
	default:
		fmt.Println("指定した拡張子がまちがっている可能性があります")
		os.Exit(1)
	}
	return err
}
