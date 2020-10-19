package main

import (
	"cvs/conversion"
	"flag"
	"fmt"
	"os"
)

var (
	extension string
	filepath  string
)

func main() {

	flag.StringVar(&extension, "e", "jpeg", "拡張子の指定")
	flag.StringVar(&filepath, "f", "", "ファイルのパスの指定")
	flag.Parse()
	err := conversion.ExtensionCheck(extension)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	err = conversion.FilepathCheck(filepath)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

}
