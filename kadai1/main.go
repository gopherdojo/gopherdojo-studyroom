package main

import (
	"fmt"
	"highpon/args"
	"os"
	"path/filepath"
)

func main() {
	var a = args.ParseArgs()
	fmt.Println(a.From)
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("path: %#v\n", path)
		return nil
	})
	fmt.Println(err)
}
