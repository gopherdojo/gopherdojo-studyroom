package main

import (
	"os"

	"github.com/gopherdojo-studyroom/kadai2/hiroya-w/imgconv"
)

func main() {
	cli := &imgconv.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run())
}
