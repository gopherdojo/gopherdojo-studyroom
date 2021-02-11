package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/shinnosuke-K/gopherdojo-studyroom/kadai3-1/shinnosuke-K/word"
)

func main() {

	num := flag.Int("n", 10, "number of questions")
	flag.Parse()

	correct := 0
	incorrect := 0

	buf := bufio.NewScanner(os.Stdin)
	w := make(chan string, 1)
	go func() {
		for buf.Scan() {
			w <- buf.Text()
		}
	}()

	rand.Seed(time.Now().Unix())

	for i := 0; i < *num; i++ {
		fmt.Print("word:")
		q := word.List[rand.Intn(len(word.List))]
		fmt.Println(q)
		fmt.Print("input:")
		select {
		case r := <-w:
			if r == q {
				fmt.Println("◯")
			} else {
				fmt.Println("×")
			}
		}
	}

	fmt.Println()
	fmt.Printf("correct:%d\n", correct)
	fmt.Printf("incorrect:%d\n", incorrect)
	fmt.Printf("rate:%.3f\n", float64(correct)/float64(*num)*100)
}
