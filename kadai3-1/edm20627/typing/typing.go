package typing

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

func Start(r io.Reader, w io.Writer, words []string, gameTime time.Duration) int {
	var score int
	timeLimit := time.After(gameTime)
	ch := input(r)
	fmt.Fprintln(w, "game start!!")

L:
	for _, word := range words {
		fmt.Fprintln(w, word)
		fmt.Fprint(w, ">")
		select {
		case answer := <-ch:
			if word == answer {
				score++
			}
		case <-timeLimit:
			fmt.Fprintf(w, "\ntime out!!\n")
			break L
		}
	}
	fmt.Fprintln(w, "score: ", score)
	return score
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}
