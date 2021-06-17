package main

import (
	"flag"
	"log"
	"os"
	"time"

	typing "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-1/Mizushima/typing"
)

func main() {
	wordsPath := "./gamedata/words.csv"

	tm := flag.Duration("limit", 20, "Time limit of the game")
	flag.Parse()

	TimeLimit := time.Second * (*tm)
	_, err := typing.Game(os.Stdin, os.Stdout, wordsPath, TimeLimit, false)
	if err != nil {
		log.Fatal(err)
	}
}
