package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// getWrodGenerator returns wordGenerator function which returns random package name of Go in type of string
func getWordGenerator() func() string {
	// word list
	words := [...]string{"archive", "bufio", "builtin", "bytes", "cmd",
		"compress", "container", "context", "crypto", "database", "debug", "embed",
		"encoding", "errors", "expvar", "flag", "fmt", "go", "hash", "html", "image",
		"index", "io", "log", "math", "mine", "net", "os", "path", "plugin", "reflect",
		"regexp", "runtime", "sort", "strconv", "strings", "sync", "syscall",
		"testing", "text", "time", "unicode", "unsafe"}

	// shuffle word list
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })

	// return function which returns a word from the word list
	i := -1
	return func() string {
		i++
		if i >= len(words) {
			i = 0
		}
		return words[i]
	}
}

// timer sleeps timesec seconds and then write to channel to notify times up
func timer(ch chan<- int, timesec time.Duration) {
	time.Sleep(timesec * time.Second)
	ch <- 1
}

// scanInput scans a word from console and write to channel
func scanInput(ch chan<- string) {
	var input string
	fmt.Scan(&input)
	ch <- input
}

func main() {
	chInput := make(chan string) // channel for scanInput
	chTimer := make(chan int)    // channel for timer

	// obtain wordGenerator which returns random package name of Go in type of string
	wordGen := getWordGenerator()

	// start timer for 10 sec
	go timer(chTimer, 10)

	score := 0
	for i := 0; i < 10; i++ {
		// get a word and print to console
		word := wordGen()
		fmt.Printf("%s\n> ", word)

		// scan a word from console
		go scanInput(chInput)

		// select channels
		select {
		case input := <-chInput:
			// check input and add score
			if word == input {
				fmt.Println("OK :)")
				score += 10
			} else {
				fmt.Println("KO :(")
			}
		case <-chTimer:
			// times up
			fmt.Println("\nTimes up!!")
			i = math.MaxInt32 - 1 // set i biggest int to break loop
		}
	}

	// output results to console
	fmt.Printf("Score: %d: ", score)
	switch {
	case score >= 80:
		fmt.Printf("Awesome :D\n")
	case score >= 50:
		fmt.Printf("OK :)\n")
	default:
		fmt.Printf("KO :(\n")
	}
}
