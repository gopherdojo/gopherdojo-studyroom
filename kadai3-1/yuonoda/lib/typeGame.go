package typeGame

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

// 別ルーチンで入力値を受け付ける
func receiveInput(s *bufio.Scanner) <-chan string {
	ch := make(chan string)
	go func() {
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

func Start(limitSeconds int, words []string, r io.Reader, w io.Writer) int {
	// 変数を宣言
	var score int
	timeLimit := time.After(time.Duration(limitSeconds) * time.Second)
	isTimedUp := false
	scanner := bufio.NewScanner(r)
	inputCh := receiveInput(scanner)

	for _, word := range words {

		// 問題を出題
		fmt.Fprintf(w, "type '%s'\n", word)

		select {

		//　入力を受けたとき
		case inputWord := <-inputCh:
			if word == inputWord {
				fmt.Fprintf(w, "correct!\n")
				score++
			} else {
				fmt.Fprintf(w, "incorrect! got \"%s\", expected \"%s\"\n", inputWord, word)
			}
			break

		// 制限時間に達したとき
		case <-timeLimit:
			fmt.Fprintln(w, "time up!")
			isTimedUp = true
			break
		}
		fmt.Fprintln(w)

		// 制限時間になったら、ループを終了
		if isTimedUp {
			break
		}
	}

	fmt.Fprintf(w, "game finished! your score is %d / %d \n", score, len(words))
	return score
}
