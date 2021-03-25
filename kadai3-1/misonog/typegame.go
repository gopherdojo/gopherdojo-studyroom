package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func run() error {
	in := input(os.Stdin)
	timelimit := time.After(10 * time.Second)

	for {
		fmt.Print(">")

		select {
		case word := <-in:
			fmt.Println(word)
		case <-timelimit:
			fmt.Println()
			fmt.Println("-----")
			fmt.Println("Finish!")
			fmt.Printf("Result: %v points.\n", 5)
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
func importWords(fileName string) ([]string, error) {
	var words []string

	f, err := os.Open(fileName)
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
