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
	fmt.Printf("%v", util.Dirwalk(targetDir))
}
