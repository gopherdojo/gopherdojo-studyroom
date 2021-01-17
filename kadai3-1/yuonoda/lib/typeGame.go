package typeGame

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func input(r io.Reader) <-chan string {
	inputChan := make(chan string, 1)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			inputChan <- s.Text()
		}
		close(inputChan)
	}()
	return inputChan
}

func Start(limitSeconds int, words []string) {
	// 変数を宣言
	var score uint
	timeLimit := time.After(time.Duration(limitSeconds) * time.Second)
	isTimedUp := false

	for _, w := range words {

		// 問題を出題
		fmt.Printf("Type '%s'\n", w)

		select {
		//　入力を受けたとき
		case inputWord := <-input(os.Stdin):
			if w == inputWord {
				fmt.Printf("correct!\n")
				score++
			} else {
				fmt.Printf("incorrect! got \"%s\", expected \"%s\"\n", inputWord, w)
			}
			break

		// 制限時間に達したとき
		case <-timeLimit:
			fmt.Println("time up!")
			isTimedUp = true
			break
		}
		fmt.Println()

		// 制限時間になったら、ループを終了
		if isTimedUp {
			break
		}
	}

	fmt.Printf("game finished! your score is %d / %d \n", score, len(words))
}
