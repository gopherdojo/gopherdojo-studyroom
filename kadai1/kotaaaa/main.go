package main

import (
	"flag"
	"fmt"

	"github.com/kotaaaa/gopherdojo-studyroom/kadai1/kotaaaa/convert" // TODO GO module
)

var targetPath string
var targetSrcExt string
var targetDstExt string

func init() {
	flag.StringVar(&targetPath, "path", "", "ファイルパス")
	flag.StringVar(&targetSrcExt, "srcExt", ".jpg", "変換前の拡張子")
	flag.StringVar(&targetDstExt, "dstExt", ".png", "変換後の拡張子")
}

func main() {
	flag.Parse()
	fmt.Println("strFlag:", targetPath)
	// fileNames := convert.GetFiles(targetPath, targetSrcExt)
	fileNames := [...]string{"owl.jpg"}
	for _, fileName := range fileNames {
		convert.Convert(fileName, targetDstExt)
	}
}
