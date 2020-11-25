package main

import (
	"flag"
	"fmt"
	"gopherdojo-studyroom/kadai1/kujiraf/converter"
	"os"
)

var cvt *converter.Converter

func init() {
	cvt = &converter.Converter{}
	cvt.Src = flag.Arg(0)
	flag.StringVar(&cvt.Dst, "out", "./output", "output directory. default is `./output`")
	flag.StringVar(&cvt.From, "from", ".jpg", "extension before convert. default is `.jpg`")
	flag.StringVar(&cvt.To, "to", ".png", "extension after converted. default is `.png`")
	flag.BoolVar(&cvt.IsDebug, "debug", false, "debug message flag. default value is false.")
}

// 終了コード
const (
	ExitCodeOK = 0
	ExitError  = 1
)

func main() {
	fmt.Println("[INFO]start conversoin.")
	flag.Parse()

	// 入力チェック
	if err := cvt.Validate(); err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR]"+err.Error())
		os.Exit(ExitError)
	}

	// 変換を実行する
	if err := cvt.DoConvert(); err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR]"+err.Error())
		os.Exit(ExitError)
	}
	fmt.Println("[INFO]end conversion.")
}
