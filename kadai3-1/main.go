package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

var t int
var wordsPath = "./word.csv"

//制限時間をオプションにする
func init() {
	flag.IntVar(&t, "t", 20, "制限時間s")
	flag.Parse()
}

func main() {
	fmt.Println("タイピングゲームが始まった！制限時間", t, "s")
	num, score := 0, 0
	timeout := time.After(time.Second * time.Duration(t))
	for sign := true; sign == true; {
		word := RandomWord()
		num++
		fmt.Println("単語NO:", num, "この英単語をタイピングしてください：", word)
		c := imp(os.Stdin)
		select {
		case right := <-c:
			if right == word {
				fmt.Println("正解です！")
				score++
			} else {
				fmt.Println("残念、不正解です。")
			}
		case <-timeout:
			fmt.Println("時間です！")
			sign = false
		}
	}
	fmt.Println("ゲーム終了！ やった単語数", num, " 時間内に正確単語数：", score)
}

//英単語をランダムに取り出す
func RandomWord() (word string) {
	words, err := readCSV(wordsPath)
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().Unix())
	num := rand.Intn(len(words))
	return words[num]
}

//入力単語を取得する
func imp(r io.Reader) <-chan string {
	wordCh := make(chan string, 1)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	wordCh <- scanner.Text()
	//close(wordCh)
	return wordCh
}

func readCSV(path string) ([]string, error) {
	var ret []string
	var row []string
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

	for {
		row, err = csvFile.Read()
		if err != nil {
			break
		}
		ret = append(ret, row...)
	}

	return ret, nil
}
