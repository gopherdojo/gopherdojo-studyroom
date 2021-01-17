package main

import (
	typeGame "github.com/yuonoda/ggopherdojo-studyroom/kadai3-1/yuonoda/lib/typeGame/lib"
	"os"
)

var words = []string{
	"peach",
	"orange",
	"apple",
	"grape",
	"pineapple",
	"mandarin",
	"lemon",
	"kiwi",
	"grapefruit",
}

func main() {
	typeGame.Start(30, words, os.Stdin, os.Stdout)
}
