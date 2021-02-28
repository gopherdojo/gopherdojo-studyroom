package main

import (
	"flag"
	"log"

	"github.com/ShunyaNagashige/imgconv/image"
)

var (
	srcFormat = flag.String("s", "jpg", "変換前ファイルの拡張子")
	dstFormat = flag.String("d", "png", "変換後ファイルの拡張子")
)

func main() {
	flag.Parse()

	s := flag.Args()

	if len(s) == 0 {
		log.Fatal("ディレクトリを指定してください。")
	}

	formats := []*string{srcFormat, dstFormat}

	for _, format := range formats {
		switch *format {
		case "png", "jpg", "gif":
		case "jpeg":
			*format = "jpg"
		default:
			log.Fatalf("拡張子%sを指定することはできません。", *format)
		}
	}

	var sources []string

	//指定したディレクトリ以下にある、拡張子srcFormatのファイルのパスを取得
	for _, dir := range s {
		var err error
		sources, err = image.Search(dir, *srcFormat)
		if err != nil {
			log.Fatalf("%#v", err)
		}
		if len(sources) == 0 {
			log.Fatalf("拡張子%sのファイルが、ディレクトリ%s以下に存在しません。", *srcFormat, dir)
		}
	}

	//画像の形式変換を行う
	for _, src := range sources {
		if err := image.Convert(src, *srcFormat, *dstFormat); err != nil {
			switch err := err.(type) {
			case *image.ConvertError:
				log.Fatalf("%#v", err)
			case *image.FileError:
				log.Fatalf("%#v", err)
			}
		}
	}
}
