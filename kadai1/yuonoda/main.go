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

func getImages(path string, fmt string) ([]os.FileInfo, error) { // TODO 何を返すべきか
	// ディレクトリ内を再帰的に探索
	var files []os.FileInfo
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
			files = append(files, info)
			return nil
		})

	// エラーがあれば、エラーを返す
	if err != nil {
		return []os.FileInfo{}, err
	}
	return files, nil
}

func main() {
	flag.Parse()
	fmt.Printf("fromFmt: %s\n", *fromFmt)
	fmt.Printf("toFmt: %s\n", *toFmt)
	fmt.Printf("path: %s\n", *path)
	fmt.Println("-------")

	files, err := getImages(*path, *fromFmt)
	for _, f := range files {
		fmt.Println(f.Name())
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
