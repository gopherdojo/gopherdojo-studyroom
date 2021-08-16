package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var  selectedDirecotry string

func returnFilePath() ([]string,error) {
	var stringPath  []string

	err := filepath.Walk(selectedDirecotry,
		func(paths string, info fs.FileInfo, err error) error {
			if filepath.Ext(paths) == ".jpg" {
				stringPath = append(stringPath,paths)
			}
			return nil
		})
	return  stringPath,err
}

func init()  {
	flag.StringVar(&selectedDirecotry, "s", "", "ディレクトリを指定")
}

func main() {
	flag.Parse()
	paths,err := returnFilePath()
	fmt.Println(paths)

	if err != nil {
		fmt.Fprintf(os.Stderr,"ディレクトリ選択をしてください。")
		os.Exit(-1)
	}
}
