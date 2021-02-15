package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/shinnosuke-K/gopherdojo-studyroom/kadai3-1/shinnosuke-K/word"
)

func main() {

	num := flag.Int("n", 10, "number of questions")
	t := flag.Int("t", 3, "answer time")
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

	bc := context.Background()

	rand.Seed(time.Now().Unix())
	for i := 0; i < *num; i++ {
		fmt.Print("word:")
		q := word.List[rand.Intn(len(word.List))]
		fmt.Println(q)
		fmt.Print("input:")

		ctx, cancel := context.WithTimeout(bc, time.Duration(*t)*time.Second)
		select {
		case r := <-w:
			if r == q {
				fmt.Println("◯")
				correct++
			} else {
				fmt.Println("×")
				incorrect++
			}
		case <-ctx.Done():
			fmt.Println("time out")
			fmt.Println("×")
			incorrect++
		}
		cancel()
	}

	fmt.Println()
	fmt.Printf("correct:%d\n", correct)
	fmt.Printf("incorrect:%d\n", incorrect)
	fmt.Printf("rate:%.3f\n", float64(correct)/float64(*num)*100)
}
