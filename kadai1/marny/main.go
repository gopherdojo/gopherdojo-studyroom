package main

import (
	"cvs/conversion"
	"flag"
	"fmt"
)

func main() {
	var (
		extension string
	)

	flag.StringVar(&extension, "e", "png", "拡張子の指定")
	flag.Parse()
	ext := conversion.ExtensionCheck(extension)
	fmt.Println(ext)
}
