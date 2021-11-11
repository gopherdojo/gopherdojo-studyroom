package main

import (
	"flag"
	"fmt"
	"time"
	"typing_game/typing"
)

var (
	timeLimit = flag.Int("time-limit", 15, "time limit(seconds)")
)

func init() {
	flag.Usage = func() {
		fmt.Printf(`Usage: -time-limit TIME_LIMIT
		Use: typing game,
		Default: 30 seconds`)
	}
	flag.PrintDefaults()
}

func main() {
	var score int

	flag.Parse()
	fmt.Printf("start typing game. time limit %d seconds\n", *timeLimit)

	timeout := time.After(time.Duration(*timeLimit) * time.Second)
	now := time.Now()

	for isTimeout := false; !isTimeout; {
		word := typing.RandomWord()
		fmt.Printf("Input: %s \n", word)
		fmt.Print("> ")
		c := typing.CreateChan()
		select {
			case res := <- c:
				if word == res {
					score++
					fmt.Println("success!")
					fmt.Printf("lapsed time: %vs \n\n", int(time.Since(now).Seconds())%60)
				} else {
					fmt.Println("failure...")
					fmt.Printf("lapsed time: %vs \n\n", int(time.Since(now).Seconds())%60)
				}
			case <- timeout:
				fmt.Println("time up")
				fmt.Println("result score: ", score)
				isTimeout = true
		}
	}
}
