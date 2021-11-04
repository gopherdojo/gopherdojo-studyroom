package main

import (
	"flag"
	"fmt"

	"github.com/kotaaaa/gopherdojo-studyroom/kadai1/kotaaaa/convert"
	"github.com/kotaaaa/gopherdojo-studyroom/kadai1/kotaaaa/search"
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
	fmt.Println("targetPath:", targetPath)
	fileNames := search.GetFiles(targetPath, targetSrcExt)
	for _, fileName := range fileNames {
		err := convert.Convert(fileName, targetDstExt, targetPath)
		if err != nil {
			fmt.Println("Error Occuerrd ", err)
		}
		fmt.Println("Converted ", fileName)
	}
}
