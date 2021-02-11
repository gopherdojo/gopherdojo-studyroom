package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/shinnosuke-K/gopherdojo-studyroom/kadai3-1/shinnosuke-K/word"
)

func main() {

	buf := bufio.NewScanner(os.Stdin)

	w := make(chan string, 1)
	go func() {
		for buf.Scan() {
			w <- buf.Text()
		}
	}()

	rand.Seed(time.Now().Unix())

	for {
		fmt.Print("word:")
		q := word.List[rand.Intn(len(word.List))]
		fmt.Println(q)
		fmt.Print("input:")
		select {
		case r := <-w:
			if r == q {
				fmt.Println("◯")
			} else {
				fmt.Println("×")
			}
		}
	}
}
