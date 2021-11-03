package convert

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// ディレクトリ以下のファイル一覧を取得します。
func GetFiles(dir string, ext string) []string {
	files, err := ioutil.ReadDir(dir)
	// ここで再帰的にファイルを取得する
	if err != nil {
		log.Fatal(err)
	}

	var arr []string
	for _, file := range files {
		name := file.Name()
		fmt.Println(name)
		if filepath.Ext(name) == ext {
			arr = append(arr, name)
		}
	}
	return arr
}

func removeFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	return nil
}

func getFileNameWithoutExt(path string) string {
	// Fixed with a nice method given by mattn-san
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

// filename をdst に変換します。
func Convert(srcFilename string, dstExt string) error {

	dstFileName := getFileNameWithoutExt(srcFilename) + dstExt
	fmt.Println("srcFileName: ", srcFilename)
	fmt.Println("dstFileName: ", dstFileName)
	// 変換対象ファイルをopen
	srcFile, err := os.Open(srcFilename)
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
	dstFile, err := os.Create(dstFileName)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// ファイル変換実行
	switch filepath.Ext(dstFileName) {
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
	removeFile(srcFilename)

	if err != nil {
		return err
	}
	return nil
}
