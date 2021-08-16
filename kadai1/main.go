package main

import (
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
)

var  selectedDirecotry string

func returnFilePath() ([]string,error) {
	var stringPath  []string

	err := filepath.Walk(selectedDirecotry,
		func(path string, info fs.FileInfo, err error) error {
			if filepath.Ext(path) == ".jpg" {
				fmt.Println(path)
				stringPath = append(stringPath,path)
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
	var path,err = returnFilePath()
	if err != nil {
		fmt.Println(path)
	}
}
