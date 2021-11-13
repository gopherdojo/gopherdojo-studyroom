package starter

import (
	"fmt"
	"os"
	"time"

	questions "github.com/kotaaaa/gopherdojo-studyroom/kadai3-1/kotaaaa/question"
)

var vc *questions.Vocab

const wordPath string = "./testdata/ielts_words.txt"
const limit_time int = 30

func Run() {

	vc = questions.ReadWords(wordPath)
	ch := make(chan string)
	var cnt int
	var input string
	timeCh := time.NewTimer(time.Duration(limit_time) * time.Second)
	for {
		idx := questions.CreateProblem(vc)
		fmt.Println("Target: ", vc.Words[idx])
		go func() {
			// Read input by user
			fmt.Scan(&input)
			ch <- input
		}()
		select {
		case <-timeCh.C:
			fmt.Println("\n======================\nTimeup! point: ", cnt, "pt\n======================") // time up process
			os.Exit(0)
		case val := <-ch:
			if val == vc.Words[idx] {
				cnt++
				fmt.Print("Correct! ", cnt, "pt ") // Correct answer
			} else {
				fmt.Print("Miss! ", cnt, "pt ") // Missed answer
			}
			fmt.Println(vc.Meanings[idx])
		}
	}
}
