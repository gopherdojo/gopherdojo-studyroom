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

func srcFileList(dir string, from *string) ([]string, error) {
	var srcFileList []string
	err := filepath.Walk(dir,
		func(srcPath string, info os.FileInfo, err error) error {
			if filepath.Ext(srcPath) == *from {
				srcFileList = append(srcFileList, srcPath)
			}
			return nil
		})
	return srcFileList, err
}

func main() {
	flag.Parse()
	dir := flag.Arg(0)

	// 変換前ファイルのリストを取得
	filelist, err := srcFileList(dir, from)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	}

	// 変換
	for _, srcPath := range filelist {
		convimgerr := convimg.Do(srcPath, convimg.Ext(*to), *rmSrc)
		if convimgerr != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", convimgerr)
		}
	}
}
