package main

import (
	"flag"
	"log"
	"strings"

	"github.com/taka521/gopherdojo-studyroom/kadai1/taka521/conv"
	"github.com/taka521/gopherdojo-studyroom/kadai1/taka521/conv/constant"
)

var from, to string

func init() {
	flag.StringVar(&from, "f", constant.ExtensionJpeg.S(), "拡張子 (from)")
	flag.StringVar(&to, "t", constant.ExtensionPng.S(), "拡張子　(to)")
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("[Error] ディレクトリは必ず指定してください。")
	}

	i := conv.HandlerInput{Dir: flag.Arg(0), From: strings.ToLower(from), To: strings.ToLower(to)}
	if err := i.Validate(); err != nil {
		log.Fatalf("[Error] %v\n", err)
	}

	if err := conv.Handle(i); err != nil {
		log.Fatalf("[Error] %v\n", err)
	}
}
