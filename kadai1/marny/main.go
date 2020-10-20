package main

import (
	"cvs/conversion"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	extension string
	imagepath string
	dirpath   string
)

func main() {

	flag.StringVar(&extension, "e", "jpeg", "拡張子の指定")
	flag.StringVar(&imagepath, "f", "", "ファイルのパスの指定")
	flag.StringVar(&dirpath, "d", "./", "保存先のパス指定")
	flag.Parse()
	err := conversion.ExtensionCheck(extension)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	err = conversion.FilepathCheck(imagepath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	f := filepath.Ext(imagepath)
	err = conversion.FileExtCheck(f)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("よくぞここまで来た")
	fmt.Println("変換中・・・")

	err = conversion.FileExtension(extension, imagepath, dirpath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
