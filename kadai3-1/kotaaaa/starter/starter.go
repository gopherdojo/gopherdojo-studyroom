package starter

import (
	"fmt"
	"time"

	"github.com/kotaaaa/gopherdojo-studyroom/kadai3-1/kotaaaa/questions"
)

type Status int

const (
	Success Status = iota
	Fail
	TimeUp
)

func Solve(vc *questions.Vocab, timeCh *time.Timer) Status {

	var input string // for input from user
	idx := questions.CreateProblem(vc)
	fmt.Println("Target: ", vc.Words[idx])
	ch := make(chan string)
	go func() {
		// Read input by user
		fmt.Scan(&input)
		ch <- input
	}()
	// each cases
	select {
	case <-timeCh.C:
		return TimeUp
	case val := <-ch: // if there is user's input
		fmt.Println("Meaning: ", vc.Meanings[idx])
		if val == vc.Words[idx] {
			return Success
		} else {
			return Fail
		}
	}
}
