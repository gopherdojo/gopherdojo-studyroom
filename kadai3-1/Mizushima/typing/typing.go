package typing

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func Game(r io.Reader, w io.Writer, words []string, t time.Duration, isTest bool) (int, error) {

	limit := time.After(t)

	rand.Seed(time.Now().UnixNano())
	var indices []int
	indices = rand.Perm(len(words))
	if isTest {
		indices = incSlice(len(words))
	}

	if !isTest {
		_, err := fmt.Fprintln(w, "> Typing game start\nPlease type the displayed word")
		if err != nil {
			return -1, err
		}

		_, err = fmt.Fprintf(w, "> Time limit is %d seconds\n", int(t.Seconds()))
		if err != nil {
			return -1, err
		}
	}

	ch := input(r)
	score := 0
	var idx int = 0
	var word string

	for {

		word = words[indices[idx]]

		_, err := fmt.Fprintf(w, "> %s\n", word)
		if err != nil {
			return -1, err
		}

		idx++

		select {
		case <-limit:
			_, err = fmt.Fprintf(w, "\nGame ends!\nThe number of correct answers is %d\n", score)
			if err != nil {
				return -1, err
			}
			return score, nil
		case ans := <-ch:
			if ans == word {
				if !isTest {
					_, err = fmt.Fprintln(w, "> しぇえか～い")
					if err != nil {
						return -1, err
					}
				}
				score++
			} else {
				if !isTest {
					_, err = fmt.Fprintln(w, "> ぶっぶー")
					if err != nil {
						return -1, err
					}
				}
			}
		}
	}
}

// input returns a channel receives string.
func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		// close(ch)
	}()
	return ch
}

// incSlice returns the integer slice starts with 0 and ends with n-1.
func incSlice(n int) []int {
	var res []int
	for i := 0; i < n; i++ {
		res = append(res, i)
	}
	return res
}
