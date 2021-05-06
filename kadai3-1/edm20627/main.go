package main

import (
	"os"
	"time"

	"github.com/edm20627/gopherdojo-studyroom/kadai3-1/edm20627/typing"
)

var words = []string{
	"go",
	"java",
	"ruby",
	"php",
	"javascript",
	"python",
	"kotlin",
	"swift",
	"c",
}

var gameTime = 20 * time.Second

func main() {
	typing.Start(os.Stdin, os.Stdout, words, gameTime)
}
