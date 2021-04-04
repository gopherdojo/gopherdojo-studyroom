package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/kimuson13/gopherdojo-studyroom/kimuson13/typing/word"
)

func main() {
	quizlist := word.MakeQuiz()
	var correct int
	inputCh := input(os.Stdin)
	timeoutCh := time.After(20 * time.Second)
	fmt.Print("You need type word is displayed as question\nTime limits is 20 seconds\n")
L:
	for _, q := range quizlist {
		fmt.Print("question:\n")
		fmt.Print(q + "\n")
		fmt.Print("input:\n")
		select {
		case r := <-inputCh:
			if r == q {
				fmt.Print("correct!\n")
				correct++
			} else {
				fmt.Print("incorrect!\n")
			}
		case <-timeoutCh:
			fmt.Print("\ntimed out!\n")
			break L
		}
	}
	fmt.Printf("Your result: %v", correct)
}

func input(r io.Reader) <-chan string {
	ch1 := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch1 <- s.Text()
		}
		close(ch1)
	}()
	return ch1
}
