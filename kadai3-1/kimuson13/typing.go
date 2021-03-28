package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/kimuson13/gopherdojo-studyroom/kimuson13/typing/word"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	quiz := rand.Perm(len(word.WordList))
	var correct int
	ch1 := input(os.Stdin)
	ch2 := time.After(20 * time.Second)
	fmt.Print("You need type word is displayed as question\nTime limits is 20 seconds\n")
	for _, q := range quiz {
		fmt.Print("question:\n")
		fmt.Print(word.WordList[q] + "\n")
		fmt.Print("input:\n")
		select {
		case r := <-ch1:
			if r == word.WordList[q] {
				fmt.Print("correct!\n")
				correct++
			} else {
				fmt.Print("incorrect!\n")
			}
		case <-ch2:
			fmt.Print("\ntimed out!\n")
			fmt.Printf("Your result: %v", correct)
			os.Exit(0)
		}
	}
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
