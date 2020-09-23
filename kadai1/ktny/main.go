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

	// println(targetDir, *from, *to)

	// targetDir配下を再帰的に画像変換する
	filepaths := util.DirWalk(targetDir)
	fmt.Printf("targetDir: %s, from: %s, to: %s\n", targetDir, *from, *to)
	for _, filepath := range filepaths {
		switch *from {
		case "jpg", "jpeg":
			if strings.HasSuffix(filepath, ".jpg") || strings.HasSuffix(filepath, ".jpeg") {
				fmt.Printf("convert %s\n", filepath)
				util.ConvertImage(filepath, *from, *to)
			}
		case "png":
			if strings.HasSuffix(filepath, ".png") {
				fmt.Printf("convert %s\n", filepath)
				util.ConvertImage(filepath, *from, *to)
			}
		}
	}
}
