package main

import (
	"os"

	"github.com/kynefuk/gopherdojo-studyroom/kadai1/cli"
)

func main() {
	cli := cli.CLI{
		OutStream: os.Stdout,
		ErrStream: os.Stderr,
	}
	os.Exit(cli.Run(os.Args))
}
