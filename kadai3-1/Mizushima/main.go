package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {

	words, err := readCSV("./words.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(words)

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

		idx := rand.Intn(len(words))
		fmt.Printf("> %s\n", words[idx])

		select {
		case <-ctx.Done():
			fmt.Println("\n終了!")
			fmt.Printf("%d問正解です!\n", score)
			return
		case <-time.After(1 * time.Second):
			if <-ch == words[idx] {
				fmt.Println("> しぇえか～い")
				score++
			} else {
				fmt.Println("> ぶっぶー")
			}
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

func readCSV(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

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