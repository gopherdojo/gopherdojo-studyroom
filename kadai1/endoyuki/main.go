package main

import (
	"flag"
	"log"
	"os"

	"endoyuki/conversion"
)

type Hoge struct {
	inputDir  string  //指定するディレクトリ
	outputDir string  //出力先ディレクトリ
	beforeExt *string //変換前の拡張子
	afterExt  *string //変換後の拡張子
}

var hoge Hoge

func main() {
	hoge.beforeExt = flag.String("b", "jpg", "before Extension")
	hoge.afterExt = flag.String("a", "png", "after Extension")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("wrong number of arguments")
	}

	for i := 0; i < len(args); i++ {
		existDir(args[i])
		args[i] = args[i] + dirPathFormating(args[i])
	}
	hoge.inputDir = args[0]
	hoge.outputDir = args[1]

	conversion.Convert(hoge.inputDir, hoge.outputDir, hoge.beforeExt, hoge.afterExt)
}

func existDir(dirName string) {
	if f, err := os.Stat(dirName); os.IsNotExist(err) || !f.IsDir() {
		log.Fatal(err)
	}
}

func dirPathFormating(dirPath string) string {
	trailChr := dirPath[len(dirPath)-1 : len(dirPath)]
	switch trailChr {
	case "/":
		return ""
	case "\\":
		return ""
	default:
		return "/"
	}
}
