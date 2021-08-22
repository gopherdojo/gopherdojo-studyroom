package main

import (
	"flag"
	"fmt"
)

type Arguments struct {
	selectedDirectory string
	selectedFileType  string
	convertedFileType string
	stringPath        []string
	isHelp            bool

	args              []string

}

func ParseArguments() (*Arguments, error) {
	var argument Arguments

	flag.StringVar(&argument.selectedDirectory, "s", "", "ディレクトリを指定")
	flag.StringVar(&argument.selectedFileType, "f", ".jpg", "変換前のファイルタイプを指定")
	flag.StringVar(&argument.convertedFileType, "cf", ".png", "変換後のファイルタイプを指定")
	flag.BoolVar(&argument.isHelp, "help", false, "display this help and exit")

	flag.Parse()

	return &argument, nil

}

func help() {
	fmt.Println(`
Usage:
 convert [options] command
	Options:
  	-s,  変換したいファイルがあるディレクトリを指定
  	-f,  変換前のファイルタイプを指定 
  	-cf, 変換後のファイルタイプを指定`,
	)
}
