package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func run(in <-chan string, w io.Writer) error {
	// in := input(os.Stdin)
	timelimit := time.After(10 * time.Second)

	words, err := importWords("word_list.txt")
	if err != nil {
		return err
	}

	typing := &Typing{Words: words}
	typing.shuffle()

	for {
		fmt.Fprintln(w, typing.Word)
		// fmt.Println(typing.Word)
		fmt.Print(">")

		select {
		case word := <-in:
			if typing.check(word) {
				fmt.Println("+1")
				fmt.Println()
				typing.plus()
			} else {
				fmt.Println("Wrong input")
				fmt.Println()
			}
			typing.shuffle()

		case <-timelimit:
			fmt.Println()
			fmt.Println("-----")
			fmt.Println("Finish!")
			fmt.Printf("Result: %v points.\n", typing.getPoint())
			return nil
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
		close(ch)
	}()
	return ch
}

// textファイルを読み込んで、出題用wordのスライスを返す
func importWords(filePath string) ([]string, error) {
	var words []string

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())

	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return words, nil
}
