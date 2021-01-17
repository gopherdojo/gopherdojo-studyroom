package typeGame

import (
	"fmt"
)

//func takeRandomWord(l []string) string {
//	log.Println("takeRandomWord")
//	rand.Seed(time.Now().UnixNano())
//	log.Println("len(l):", len(l))
//	i := rand.Intn(len(l))
//	log.Println("i:", i)
//	return l[i]
//}

func Start() {
	var input string
	var score uint
	words := []string{
		"peach",
		"orange",
		"apple",
		"grape",
		"pineapple",
		"mandarin",
		"lemon",
		"kiwi",
		"grapefruit",
	}

	for {
		// 単語を表示して、入力を受ける
		//word := takeRandomWord(words)
		word := words[0]
		words = words[1:]
		fmt.Printf("Type '%s'\n", word)
		fmt.Scan(&input)
		if input == word {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Printf("Incorrect! got \"%s\", expected \"%s\"\n", input, word)
		}
		fmt.Println("score:", score)
		fmt.Println("")
	}
}
