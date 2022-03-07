package main

import (
	"flag"
	"watcher041/convExt"
)

var (
	beforeExt = flag.String("beforeExt", "jpg", "変換前のオプション")
	afterExt  = flag.String("afterExt", "png", "変換後のオプション")
)

func init() {

	// オプションで指定した値をここで変数に代入する
	flag.Parse()

}

func main() {

	// 画像の拡張子を変換する
	convExt.ConvExt(*beforeExt, *afterExt)
}
