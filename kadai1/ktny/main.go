package main

import (
	"flag"
	"fmt"
	"kadai1/ktny/util"
	"os"
	"strings"
)

func main() {
	var from = flag.String("from", "jpg", "変換前の画像ファイルの拡張子")
	var to = flag.String("to", "png", "変換後の画像ファイルの拡張子")

	flag.Parse()

	targetDir := flag.Arg(0)
	if targetDir == "" {
		fmt.Println("ディレクトリが指定されていません")
		os.Exit(1)
	}
	targetDir = strings.TrimRight(targetDir, "/")

	// targetDir配下を再帰的に画像変換する
	// convertImageDir(targetDir, *from, *to)
	println(targetDir, *from, *to)

	util.ConvertImage("./sample.jpg", *from, *to)

	// filepaths := util.DirWalk(targetDir)
	// for _, filepath := range filepaths {
	// 	extension := util.GetExtension(filepath)
	// 	if util.Contains(util.ImageExtensions, extension) {
	// 		fmt.Printf("convert %s from %s to %s\n", filepath, *from, *to)
	// 		util.ConvertImage(filepath, *from, *to)
	// 	}
	// }
}
