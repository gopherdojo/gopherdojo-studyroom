package main

import (
	_ "embed"
	"flag"
	"log"
	"os"
	"strings"
	"time"

	typing "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-1/Mizushima/typing"
)

//go:embed typing/gamedata/words.csv
var words string

func main() {

	wordsSlice := strings.Split(words, ",")
	tm := flag.Duration("limit", 20, "Time limit of the game")
	flag.Parse()

	TimeLimit := time.Second * (*tm)
	_, err := typing.Game(os.Stdin, os.Stdout, wordsSlice, TimeLimit, false)
	if err != nil {
		log.Fatal(err)
	}
}
