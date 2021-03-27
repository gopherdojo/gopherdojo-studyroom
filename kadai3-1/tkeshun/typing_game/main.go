package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	"typing_game/word"
)

var jp bool
var qty *int
var tl *int

func init() {

	qty = flag.Int("qty", 100, "対象個数を選ぶ")
	if *qty > 900 {
		*qty = 900
	}
	tl = flag.Int("tl", 20, "制限時間(秒)指定")
}

/*標準入力*/
func getInput(input chan string) {
	in := bufio.NewReader(os.Stdin)
	result, err := in.ReadString('\n')
	if err != nil {
		in = nil
	}
	input <- result
}

func main() {
	//flagのパース
	flag.Parse()
	setumei := "NAWL(NEW ACADEMIC WORD LIST)から単語を選んで出力する\n出力された英単語をタイピングしてEnterを押すことで問題に回答できます\n"
	fmt.Printf(setumei)

	//単語リストの取得
	wordList, err := word.GetWordList(*qty)
	if err != nil {
		log.Fatal(err)
	}

	seikai := 0
	index_length := len(wordList)
	max_num := 0
	//問題出題
	f := func() string {
		index := rand.Intn(index_length)
		fmt.Printf("%d %s\n>> ", max_num+1, wordList[index][0])
		return wordList[index][0]
	}

	ctx := context.Background()
	//d := 20000 * time.Millisecond
	d := time.Duration(*tl) * time.Second
	timer, cancel := context.WithTimeout(ctx, d)
	defer cancel()

LOOP:
	for {
		CorrectAnswer := f()
		max_num++
		input := make(chan string, 1)
		ans := ""
		go getInput(input)

		select {

		case <-timer.Done():
			if ans == "" {
				//入力前のタイムアウトに対応
				fmt.Printf("\n")
			}
			fmt.Printf("%d/%d 正解\n", seikai, max_num)
			break LOOP
		case s := <-input:
			ans = s
			if max_num%10 == 0 {
				rand.Seed(time.Now().UnixNano())
			}
			if CorrectAnswer == ans {
				fmt.Print("○\n")
				seikai++
			} else {
				fmt.Print("X\n")
			}
		}
	}
}
