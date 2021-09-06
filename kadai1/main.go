package main

import (
	"flag"
	"image-convert/convert"
	"log"
)

type cmdArgs struct {
	from string
	to   string
	dir  string
}

func parseArgs() *cmdArgs {
	var cmd cmdArgs
	flag.StringVar(&cmd.from, "from", "jpg", "from")
	flag.StringVar(&cmd.to, "to", "png", "to")
	flag.StringVar(&cmd.dir, "dir", "./", "directory")
	flag.Parse()
	return &cmd
}

func main() {
	var args = parseArgs()
	var convertList, err = convert.GetSelectedExtensionPath(args.from, args.dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range convertList {
		err := convert.ConvertImage(v, args.to)
		if err != nil {
			log.Fatal(err)
		}
	}
}
