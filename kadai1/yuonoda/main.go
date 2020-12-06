package main

import (
	"bytes"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
)

func main() {

	imageBytes, err := ioutil.ReadFile("./kadai1/gopher.jpg")
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	if err != png.Encode(buf, img) {
		log.Fatal(err)
	}

	//log.Println(reflect.TypeOf(buf))
	ioutil.WriteFile("gopher.png", buf.Bytes(), 0644)

}
