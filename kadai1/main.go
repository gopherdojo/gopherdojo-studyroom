package main

import (
	"image-convert/args"
	"image-convert/convert"
	"strings"
)

func main() {
	var args = args.ParseArgs()
	var convertList = convert.GetSelectedExtensionPath(args.From, args.Dir)
	for _, v := range convertList {
		convert.ConvertImage(strings.Join(v[:len(v)-1], "."), args.From, args.To)
	}
}
