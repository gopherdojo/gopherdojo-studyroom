package main

import (
	"flag"
	"fmt"
	"kadai1/ktny/util"
	"os"
	"strings"
)

func main() {
	var from = flag.String("from", "jpg", "Image file extension before conversion")
	var to = flag.String("to", "png", "Image file extension after conversion")

	flag.Parse()

	targetDir := flag.Arg(0)
	if targetDir == "" {
		fmt.Println("Directory is not defined.")
		os.Exit(1)
	}
	targetDir = strings.TrimRight(targetDir, "/")

	// targetDir配下を再帰的に画像変換する
	filepaths := util.DirWalk(targetDir)
	fmt.Printf("targetDir: %s, from: %s, to: %s\n", targetDir, *from, *to)
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
