package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {

	words := []string{"ahoaho", "bakabaka", "unkounko"}
	rand.Seed(time.Now().UnixNano())

	bc := context.Background()
	t := 20 * time.Second
	ctx, cancel := context.WithTimeout(bc, t)
	defer cancel()

	fmt.Println("> タイピングゲームを始めましゅ")
	fmt.Println("> 英単語が出てきますので、同じ単語をタイプしてくだしゃい!")
	fmt.Println("> 制限時間は20秒です")

	ch := input(os.Stdin)
	score := 0

	for {

		idx := rand.Intn(3)
		fmt.Printf("> %s\n", words[idx])

		select {
		case <-time.After(1 * time.Second):
			if <-ch == words[idx] {
				fmt.Println("> しぇえか～い")
				score++
			} else {
				fmt.Println("> ぶっぶー")
			}
		case <-ctx.Done():
			fmt.Println("\n終了!")
			fmt.Printf("%d問正解です!\n", score)
			return
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
