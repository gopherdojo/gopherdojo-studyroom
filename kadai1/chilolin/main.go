package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/chilolin/image-conversion/convertor"
)

var srcFormat string
var destFormat string

func init() {
	flag.StringVar(&srcFormat, "src", "jpg", "「'jpg', 'png', 'gif'」の中から変換前の画像形式を指定する")
	flag.StringVar(&destFormat, "dest", "png", "「'jpg', 'png', 'gif'」の中から変換後の画像形式を指定する")
}

func main() {
	flag.Parse()

	err := convertor.Do(flag.Args()[0], srcFormat, destFormat)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
