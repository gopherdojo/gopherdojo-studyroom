package main

import typeGame "github.com/yuonoda/ggopherdojo-studyroom/kadai3-1/yuonoda/lib/typeGame/lib"

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
	typeGame.Start(10, words)
}
