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
)

var fromFmt = flag.String("from", "jpg", "Specify a format of your original image")
var toFmt = flag.String("to", "png", "Specify a format which you want to convert the image to")

func main() {
	flag.Parse()
	fmt.Printf("fromFmt: %s\n", *fromFmt)
	fmt.Printf("toFmt: %s\n", *toFmt)

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
