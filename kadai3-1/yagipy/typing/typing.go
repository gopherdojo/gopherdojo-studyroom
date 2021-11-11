package typing

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func RandomWord() string {
	words := []string {
		"reason",
		"secret",
		"gimlet",
		"escape",
		"galaxy",
		"breeze",
		"beetle",
		"allure",
		"velvet",
	}

	rand.Seed(time.Now().UnixNano())

	return words[rand.Intn(len(words))]
}

func CreateChan() <-chan string {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	c := make(chan string, 1)
	c <- text
	return c
}
