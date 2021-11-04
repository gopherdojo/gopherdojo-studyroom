package convert

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func removeFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	return nil
}

func getFilePathFromBase(path string) string {
	// Fixed with a nice method given by mattn-san
	return path[:len(path)-len(filepath.Ext(path))]
}

// filename をdst に変換します。
func Convert(srcFilename string, dstExt string, path string) error {

	dstFileName := getFilePathFromBase(srcFilename) + dstExt
	// 変換対象ファイルをopen
	srcFile, err := os.Open(path + srcFilename)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 変換対象ファイルの画像読み込み
	img, _, err := image.Decode(srcFile)
	if err != nil {
		return err
	}

	// 変換後ファイル作成
	dstFile, err := os.Create(path + dstFileName)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// ファイル変換実行
	switch filepath.Ext(path + dstFileName) {
	case ".gif":
		err = gif.Encode(dstFile, img, nil)
	case ".png":
		err = png.Encode(dstFile, img)
	case ".jpg", "jpeg":
		err = jpeg.Encode(dstFile, img, nil)
	}
	if err != nil {
		return err
	}

	// srcファイル削除
	err = removeFile(path + srcFilename)

	if err != nil {
		return err
	}
	return nil
}
