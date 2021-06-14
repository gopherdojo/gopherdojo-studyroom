package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {

	words := []string{ "ahoaho", "bakabaka", "unkounko" }
	rand.Seed(time.Now().UnixNano())
	ch := input(os.Stdin)
	for {
		idx := rand.Intn(3)
		fmt.Printf("> %s\n", words[idx])
		if <-ch == words[idx] {
			fmt.Println("> しぇいかい")
		} else {
			fmt.Println("> ぶっぶー")
		}
	}
}

func input(r io.Reader) <-chan string {
	// TODO: チャネルを作る
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			// TODO: チャネルに読み込んだ文字列を送る
			str := s.Text()
			ch <- str
		}
		// TODO: チャネルを閉じる
		close(ch)
	}()
	// TODO: チャネルを返す
	return ch
}
