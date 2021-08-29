package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type cmdArgs struct {
	from string
	to   string
	dir  string
}

var cmd cmdArgs

func init() {
	flag.StringVar(&cmd.from, "from", "jpg", "from")
	flag.StringVar(&cmd.to, "to", "png", "to")
	flag.StringVar(&cmd.dir, "dir", "./", "directory")
}

func main() {
	flag.Parse()
	fmt.Println(cmd.from, cmd.to, cmd.dir)

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("path: %#v\n", path)
		return nil
	})
	fmt.Println(err)
}
