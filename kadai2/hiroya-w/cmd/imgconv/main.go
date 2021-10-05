package main

import (
	"os"

	imgconv "github.com/Hiroya-W/gopherdojo-studyroom/kadai2/hiroya-w"
)

func main() {
	cli := &imgconv.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	os.Exit(cli.Run())
}
