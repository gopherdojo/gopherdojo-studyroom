package main

import (
    "flag"
    "fmt"
    "typing/typing"
    "github.com/fatih/color"
    "time"
)

// initialize flag
var (
    tim = flag.Int("time", 30, "time limit")
)

func init() {
    flag.Usage = func() {
        fmt.Printf(`Usage: -time TIME LIMIT
        Use: typing game (only English).
        Default: time 30 seconds.
        `)
        flag.PrintDefaults()
    }
}

func main() {
    flag.Parse()
    tim := *tim
    fmt.Printf("start typing game. time limit %d second\n", tim)
    var score int
    var judge = true

    timeout := time.After(time.Duration(tim) * time.Second)
    now := time.Now()
    for judge {
        word := typing.RandomWord()
        fmt.Printf("右の文字を入力せよ : %s  \n", word)
        c := typing.CreateChan(word)
        select {
            case res := <- c:
                if word == res {
                    score++
                    fmt.Printf("経過: %vs\n", int(time.Since(now).Seconds())%60)
                } else {
                    color.Red("fail")
                    fmt.Printf("経過: %vs\n", int(time.Since(now).Seconds())%60)
                }
            case <- timeout:
                fmt.Println("Time up")
                fmt.Println("result score: ", score)
                judge = false
        }
    }
}
