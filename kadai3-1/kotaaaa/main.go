package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kotaaaa/gopherdojo-studyroom/kadai3-1/kotaaaa/questions"
	"github.com/kotaaaa/gopherdojo-studyroom/kadai3-1/kotaaaa/starter"
)

const wordPath string = "./testdata/ielts_words.txt" // Vocabulary file (Format > "word:meaning")
const limit_time int = 30                            // Limit time

var vc *questions.Vocab

func main() {
	vc = questions.ReadWords(wordPath)                               // read vocab file
	timeCh := time.NewTimer(time.Duration(limit_time) * time.Second) // start timer

	var cnt int

	for {
		c := starter.Solve(vc, timeCh)
		switch c {
		case starter.Success:
			cnt++
			fmt.Print("Correct! ", cnt, "pt \n") // Correct answer
		case starter.Fail:
			fmt.Print("Miss! ", cnt, "pt ") // Missed answer
		case starter.TimeUp:
			fmt.Println("\n======================\nTimes up! point: ", cnt, "pt\n======================") // time up process
			os.Exit(0)
		}
	}
}
