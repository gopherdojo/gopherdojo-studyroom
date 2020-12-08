package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
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

func decodeImage(buf *bytes.Reader, fmt string) (image.Image, error) {
	var img image.Image
	var err error
	switch fmt {
	case "jpg":
		img, err = jpeg.Decode(buf)
		break
	case "png":
		img, err = png.Decode(buf)
		break
	case "gif":
		img, err = gif.Decode(buf)
	default:
		err = errors.New("decode format couldn't be found")
	}
	return img, err
}

func main() {
	flag.Parse()
	fmt.Printf("fromFmt: %s\n", *fromFmt)
	fmt.Printf("toFmt: %s\n", *toFmt)
	fmt.Printf("path: %s\n", *path)
	fmt.Println("-------")

	// 画像の一覧を取得
	pathes, err := getImagePathes(*path, *fromFmt)
	if err != nil {
		log.Fatal(err)
	}
	for _, path := range pathes {
		fmt.Println(path)
	}

	// 元画像を読み込み
	path := pathes[0]
	log.Printf("%s\n", path)
	imageBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// 画像をデコード
	buffer := bytes.NewReader(imageBytes)
	img, err := decodeImage(buffer, *fromFmt)
	if err != nil {
		log.Fatal(err)
	}

	// 変換
	nameNoExt := strings.TrimSuffix(path, filepath.Ext(path))
	newBuf := new(bytes.Buffer)
	switch *toFmt {
	case "png":
		if err != png.Encode(newBuf, img) {
			log.Fatal(err)
		}
		break
	case "gif":
		if err != gif.Encode(newBuf, img, nil) {
			log.Fatal(err)
		}
		break
	default:
		log.Fatal("You cannot convert to the specified format")
	}

	// ファイル出力
	newName := nameNoExt + "." + *toFmt
	log.Println(newName)
	ioutil.WriteFile(newName, newBuf.Bytes(), 0644)

}
