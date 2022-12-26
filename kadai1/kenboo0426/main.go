package main

import (
	"flag"
	"gopherdojo-studyroom/kadai1/kenboo0426/imgconv"
	"log"
)

func main() {
	dirpath := flag.String("path", "/", "パスを指定する")
	from := flag.String("from", "jpg", "変換前の形式を指定する")
	to := flag.String("to", "png", "変換後の形式を指定する")
	flag.Parse()

	if dirpath == nil {
		log.Fatal("パスを指定する")
	}

	imgconv.Cmd(*dirpath, *from, *to)
}
