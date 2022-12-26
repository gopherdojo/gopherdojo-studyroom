package main

import (
	"flag"
	"gopherdojo-studyroom/kadai1/kenboo0426/imgconv"
	"log"
)

func main() {
	dirpath := flag.String("path", "/", "Enter the path")
	from := flag.String("from", "jpg", "Enter the img extension before conversion")
	to := flag.String("to", "png", "Enter the img extension after conversion")
	flag.Parse()

	if dirpath == nil {
		log.Fatal("パスを指定する")
	}

	imgconv.ConvertExtensions(*dirpath, *from, *to)
}
