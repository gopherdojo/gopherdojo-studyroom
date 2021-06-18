package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var t int

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
	words := []string{"banana", "apple", "milk", "fruits", "cat", "car", "elephant", "unbralla", "interface", "tissues", "format"}
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
