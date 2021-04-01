package main

import (
	"flag"
	"fmt"
	"os"

	"hsaki/convert"
)

var (
	srcDir = flag.String("src", ".", "変換前画像があるディレクトリ(相対パス)")
	dstDir = flag.String("dst", "./result", "変換後画像を配置するディレクトリ")
	bExt   = flag.String("b", "jpg", "変換前の画像拡張子")
	aExt   = flag.String("a", "png", "変換後の画像拡張子")
)

func main() {
	flag.Parse()
	cvt, err := convert.NewConverter(*srcDir, *dstDir, *bExt, *aExt)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = cvt.Do()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	//fmt.Println(os.Args)

	/*
		fmt.Println("src : ", *srcDir)
		fmt.Println("dst : ", *dstDir)
		fmt.Println("b : ", *bExt)
		fmt.Println("a : ", *aExt)
	*/
}
