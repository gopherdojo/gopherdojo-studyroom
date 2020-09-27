package main

import (
	"flag"
	"fmt"
	"kadai1/ktny/util"
	"os"
	"strings"
)

var (
	from string
	to   string
)

func init() {
	flag.StringVar(&from, "from", util.JPG, "from ext. support jpg, jpeg, png, gif.")
	flag.StringVar(&to, "to", util.PNG, "to ext. support jpg, jpeg, png, gif.")
}

func main() {
	flag.Parse()

	// ディレクトリが指定されていない場合は終了する
	targetDir := flag.Arg(0)
	if targetDir == "" {
		fmt.Println("[Error]Directory is not defined.")
		os.Exit(1)
	}
	targetDir = strings.TrimRight(targetDir, "/")

	// 指定された変換前後の拡張子が同じ場合は終了する
	if from == to {
		fmt.Println("[Error]from and to extension is same.")
		os.Exit(1)
	}

	// 指定された拡張子がサポートされていない場合は終了する
	if !util.IsSupportExt(from) {
		fmt.Printf("[Error]%s is not supported ext.", from)
	} else if !util.IsSupportExt(to) {
		fmt.Printf("[Error]%s is not supported ext.", to)
	}

	fmt.Printf("[Info]from=%s, to=%s, targetDir=%s\n", from, to, targetDir)

	// targetDir配下のファイルパスを取得する
	filepaths, err := util.DirWalk(targetDir)
	if err != nil {
		fmt.Printf("[Error]%s", err)
		os.Exit(1)
	}

	// 指定された拡張子の画像ファイルを変換する
	for _, filepath := range filepaths {
		if util.CanConvertExt(from, filepath) {
			fmt.Printf("[Info]convert %s\n", filepath)
			if err := util.ConvertImage(filepath, from, to); err != nil {
				fmt.Printf("[Error]%s", err)
			}
		} else {
			fmt.Printf("[Warn]cannnot convert %s. It is not %s file.\n", filepath, from)
		}
	}
}
