package main

import (
	"flag"
	"fmt"
	"kadai1/ktny/util"
	"os"
	"strings"
)

func main() {
	var from = flag.String("from", "jpg", "extension before conversion")
	var to = flag.String("to", "png", "extension after conversion")

	flag.Parse()

	// ディレクトリが指定されていない場合は終了する
	targetDir := flag.Arg(0)
	if targetDir == "" {
		fmt.Println("[Error]Directory is not defined.")
		os.Exit(1)
	}
	targetDir = strings.TrimRight(targetDir, "/")

	// 指定した変換前後の拡張子が同じ場合は終了する
	if *from == *to {
		fmt.Println("[Error]from and to extension is same.")
		os.Exit(1)
	}

	// targetDir配下を再帰的に画像変換する
	filepaths := util.DirWalk(targetDir)
	fmt.Printf("[Info]from=%s, to=%s, targetDir=%s\n", *from, *to, targetDir)
	for _, filepath := range filepaths {
		switch *from {
		case "jpg", "jpeg":
			if strings.HasSuffix(filepath, ".jpg") || strings.HasSuffix(filepath, ".jpeg") {
				fmt.Printf("convert %s\n", filepath)
				util.ConvertImage(filepath, *from, *to)
			} else {
				fmt.Printf("[Warn]%s was not converted. It is %s file.\n", filepath, *from)
			}
		case "png":
			if strings.HasSuffix(filepath, ".png") {
				fmt.Printf("convert %s\n", filepath)
				util.ConvertImage(filepath, *from, *to)
			} else {
				fmt.Printf("[Warn]%s was not converted. It is %s file.\n", filepath, *from)
			}
		}
	}
}
