package main

import (
	"fmt"
	"highpon/args"
	"highpon/convert"
	"strings"
)

func main() {
	var a = args.ParseArgs()
	fmt.Println(a.From)
	fmt.Println(convert.GetSelectedExtensionPath(a.From, a.Dir))
	var b = convert.GetSelectedExtensionPath(a.From, a.Dir)
	for _, v := range b {
		fmt.Println(strings.Join(v[:len(v)-1], "."))
	}

}
