package main

import (
	"bytes"
	"flag"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var fromFmt = flag.String("from", "jpg", "Specify a format of your original image")
var toFmt = flag.String("to", "png", "Specify a format which you want to convert the image to")
var path = flag.String("path", ".", "Path to a directory in which images will be converted recursively")

func getImagePathes(path string, fmt string) ([]string, error) {
	// ディレクトリ内を再帰的に探索
	var pathes []string
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			// 対象の拡張子か査定
			ext := filepath.Ext(path)[1:]
			if ext != fmt { // TODO jpgの表記
				return nil
			}
			pathes = append(pathes, path)
			return nil
		})

	// エラーがあれば、エラーを返す
	if err != nil {
		return []string{}, err
	}
	return pathes, nil
}

func main() {
	flag.Parse()
	fmt.Printf("fromFmt: %s\n", *fromFmt)
	fmt.Printf("toFmt: %s\n", *toFmt)
	fmt.Printf("path: %s\n", *path)
	fmt.Println("-------")

	// 画像の一覧を取得
	pathes, err := getImagePathes(*path, *fromFmt)
	for _, path := range pathes {
		fmt.Println(path)
	}

	return
	// 元画像を読み込み
	imageBytes, err := ioutil.ReadFile("./gopher.jpg")
	if err != nil {
		log.Fatal(err)
	}
	buffer := bytes.NewReader(imageBytes)
	jpgImg, err := jpeg.Decode(buffer)
	if err != nil {
		log.Fatal(err)
	}

	// 変換
	switch *toFmt {
	case "png":
		newBuf := new(bytes.Buffer)
		if err != png.Encode(newBuf, jpgImg) {
			log.Fatal(err)
		}
		// ファイル出力
		ioutil.WriteFile("gopher.png", newBuf.Bytes(), 0644)
		break
	case "gif":
		newBuf := new(bytes.Buffer)
		if err != gif.Encode(newBuf, jpgImg, nil) {
			log.Fatal(err)
		}
		// ファイル出力
		ioutil.WriteFile("gopher.gif", newBuf.Bytes(), 0644)
		break
	default:
		log.Fatal("You cannot convert to the specified format")
	}

}
