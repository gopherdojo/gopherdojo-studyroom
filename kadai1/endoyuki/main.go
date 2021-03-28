package main

import (
	"flag"
	"fmt"
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
		err := existDir(args[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	hoge.inputDir = args[0]
	hoge.outputDir = args[1]

	err := conversion.Convert(hoge.inputDir, hoge.outputDir, hoge.beforeExt, hoge.afterExt)
	if err != nil {
		log.Fatal(err)
	}
}

func existDir(dirName string) error {
	f, err := os.Stat(dirName)
	if err != nil {
		return err
	}
	if !f.IsDir() {
		log.Fatal(fmt.Errorf("directory %v not exist", dirName))
	}
	return nil
}
