package main

import (
	"flag"
	"fmt"
	"gopherdojo-studyroom/kadai1/okamotchan/imgconv"
	"os"
)

const CodeFailure = 1

func setArgs() (map[string]string, error) {
	from := flag.String("from", "jpg", "変更前の拡張子")
	to := flag.String("to", "png", "変更後の拡張子")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		return nil, fmt.Errorf("引数の数が正しくありません")
	}
	dir := args[0]

	// mapリテラル
	return map[string]string{"from": *from, "to": *to, "dir": dir}, nil
}

func printUsage() {
	fmt.Println("使用方法:")
	fmt.Println("  image_conversion [-from=<ext>] [-to=<ext>] target_directory")
	fmt.Println("引数:")
	fmt.Println("  -from=<ext> 変換前の拡張子", imgconv.ValidImageExt, "(default: jpg)")
	fmt.Println("  -to=<ext>   変換後の拡張子", imgconv.ValidImageExt, "(default: png)")
}

func main() {
	args, err := setArgs()
	if err != nil {
		printUsage()
		os.Exit(CodeFailure)
	}

	err = imgconv.Do(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(CodeFailure)
	}

	fmt.Println("変換が完了しました")
}
