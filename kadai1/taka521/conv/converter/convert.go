package converter

import (
	"fmt"
	"image"
	"os"

	"github.com/taka521/gopherdojo-studyroom/kadai1/taka521/conv/constant"
)

// Convert は指定された画像ファイルを、指定された拡張子の画像ファイルに変換します。
func Convert(filePath string, to constant.Extension) error {
	// 変換元ファイルオープン
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("ファイルが開けませんでした。[ファイルパス = %v]: %w", filePath, err)
	}
	defer file.Close()

	// 画像として取り扱えなければ処理終了
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// 出力ファイル名を決定
	fPath, oDir, _ := convertedPath(filePath, to)
	checkDir(oDir)

	dest, err := os.Create(fPath)
	if err != nil {
		return err
	}

	// 画像変換を実行
	return encoders[to](dest, img)
}
