package main

import (
	"flag"
	"fmt"

	"github.com/kotaaaa/gopherdojo-studyroom/kadai1/kotaaaa/convert"
	"github.com/kotaaaa/gopherdojo-studyroom/kadai1/kotaaaa/search"
	"github.com/kotaaaa/gopherdojo-studyroom/kadai1/kotaaaa/validator"
)

var targetPath string
var targetSrcExt string
var targetDstExt string

func init() {
	flag.StringVar(&targetPath, "path", "", "File path")
	flag.StringVar(&targetSrcExt, "srcExt", ".jpg", "source file extention")
	flag.StringVar(&targetDstExt, "dstExt", ".png", "destination file extention")

}

func main() {
	flag.Parse()
	validator.ValidateArgs(targetPath, targetSrcExt, targetDstExt)
	fmt.Println("targetPath:", targetPath)
	targetPath = targetPath + "/"
	// Get file list (Relative path from basePath)
	fileNames := search.GetFiles(targetPath, targetSrcExt)

	for _, fileName := range fileNames {
		fileInfo := convert.NewFileInfo(fileName, targetDstExt, targetPath)
		// 変換処理
		err := fileInfo.Convert()
		if err != nil {
			fmt.Println("Error Occuerrd ", err)
		}
		fmt.Println("Converted ", fileName)
	}
}
