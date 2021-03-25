package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"
	"typing_game/word"
)

var jp bool
var qty *int
var tl *float64

func init() {

	qty = flag.Int("qty", 100, "対象個数を選ぶ")
	if *qty > 900 {
		*qty = 900
	}
	tl = flag.Float64("tl", 120, "制限時間(秒)を指定")
}

func main() {
	flag.Parse()
	setumei := "NAWL(NEW ACADEMIC WORD LIST)から単語を選んで出力する\n出力された英単語をタイピングしてEnterを押すことで問題に回答できます\n"
	fmt.Printf(setumei)
	//単語リストの取得

	wordList, err := word.GetWordList(*qty)
	//問題出題
	if err != nil {
		log.Fatal(err)
	}
	seikai := 0

	start := time.Now()
	index_length := len(wordList)
	max_num := 0
	for {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(rand.Intn(index_length))

		var s string
		fmt.Printf("%d %s\n>> ", max_num+1, wordList[index][0])
		_ , err = fmt.Scanf("%s", &s)
		if err != nil {
			log.Fatal(err)
		}
		if wordList[index][0] == s {
			fmt.Print("○")
			seikai++
		} else {
			fmt.Print("X")
		}
		end := time.Now()
		nokori := *tl - (end.Sub(start)).Seconds()
		if nokori > 0 {
			fmt.Printf(" 残り時間%d(sec)\n", int(nokori))
		}else{
			fmt.Printf(" 残り時間%d(sec)\n", 0)
		}

		max_num++

		if nokori < 0 {
			fmt.Println(" Time UP")
			fmt.Printf("%d / %d問正解\n", seikai, max_num)
			break
		}

	}

}
