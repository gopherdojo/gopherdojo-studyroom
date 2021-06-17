package typing

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func Game(r io.Reader, w io.Writer, wordsPath string, t time.Duration, isTest bool) (int, error) {

	words, err := readCSV(wordsPath)
	if err != nil {
		return -1, err
	}

	limit := time.After(t)

	rand.Seed(time.Now().UnixNano())
	var indices []int
	if !isTest {
		indices = rand.Perm(len(words))
	} else {
		indices = incSlice(len(words))
	}

	if !isTest {
		_, err = fmt.Fprintln(w, "> Typing game start\nPlease type the displayed word")
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
		
		_, err = fmt.Fprintf(w, "> %s\n", word)
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
		case  ans := <-ch:
			// fmt.Printf("ans: %s\n", ans)
			// fmt.Printf("word: %s\n", word)
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

func readCSV(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	csvFile := csv.NewReader(file)
	csvFile.TrimLeadingSpace = true

	var ret []string
	var row []string

	for {
		row, err = csvFile.Read()
		if err != nil {
			break
		}

		ret = append(ret, row...)
	}

	return ret, nil
}

func incSlice(n int) []int {
	var res []int
	for i := 0; i < n; i++ {
		res = append(res, i)
	}
	return res
}