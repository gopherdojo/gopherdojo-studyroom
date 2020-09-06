package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"kadai1/convimg"
)

var (
	from  *string = flag.String("from", ".jpg", "ext before convarsion")
	to    *string = flag.String("to", ".png", "ext after convarsion")
	rmSrc *bool   = flag.Bool("r", false, "remove original file")
)

func main() {
	flag.Parse()

	// fmt.Println(convimg.Ext(*from))
	// fmt.Println(convimg.Ext(*to))
	// fmt.Println(*rmSrc)

	dir := flag.Arg(0)
	// fmt.Println(dir)

	err := filepath.Walk(dir,
		func(srcPath string, info os.FileInfo, err error) error {
			if filepath.Ext(srcPath) == *from {
				// fmt.Println(srcPath)
				convimgerr := convimg.Do(srcPath, convimg.Ext(*to), *rmSrc)
				if convimgerr != nil {
					fmt.Fprintln(os.Stderr, "ERROR:", convimgerr)
				}
			}
			return nil
		})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}
