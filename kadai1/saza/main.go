package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := os.Args[1]
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return err
		}

		fmt.Println(path)

		return err
	})

	if err != nil {
		panic(err)
	}
}