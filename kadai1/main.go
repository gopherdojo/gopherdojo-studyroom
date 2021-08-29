package main

import (
	"flag"
	"fmt"
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
}
