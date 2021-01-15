package lib

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func takeRandomWord(l []string) string {
	log.Println("takeRandomWord")
	rand.Seed(time.Now().UnixNano())
	log.Println("len(l):", len(l))
	i := rand.Intn(len(l))
	log.Println("i:", i)
	return l[i]
}

func Start() {
	var words = []string{
		"apple",
		"orange",
		"peach",
	}

	var input string

	for {
		// 単語を表示して、入力を受ける
		word := takeRandomWord(words)
		fmt.Printf("Type '%s'\n", word)
		fmt.Scan(&input)
		fmt.Println(input)
	}
}
