//画像変換を行うプログラム
package main

import (
	"flag"
	"fmt"
	"image_convert/conversion"
	"image_convert/findFile"
	"log"
	"path/filepath"
)

type fileFormat *string

var input_fmt fileFormat
var output_fmt fileFormat

func init() {
	input_fmt = flag.String("i_fmt", "jpeg", "画像の入力フォーマットを指定")
	output_fmt = flag.String("o_fmt", "png", "画像の出力フォーマットを指定")
}

func main() {
	//フラグの読み込み
	flag.Parse()
	args := flag.Args()
	if len(args) > 2 {
		fmt.Println("Only one directory can be specified")
	}

	//ディレクトリ内の対象ファイル一覧を取得
	f_list := findFile.Search(flag.Arg(0), *input_fmt)

	//画像変換
	for _, srcPath := range f_list {
		fmt.Println(srcPath, " >> ", srcPath[:len(srcPath)-len(filepath.Ext(srcPath))]+"."+*output_fmt)
		err := conversion.Convert(srcPath, *output_fmt)
		if err != nil {
			log.Fatal(err)
		}
	}

}
