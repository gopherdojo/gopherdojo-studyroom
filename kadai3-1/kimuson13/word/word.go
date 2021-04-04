package word

import (
	"math/rand"
	"time"
)

var wordList = []string{
	"cat",
	"dog",
	"rabbit",
	"gopher",
	"home",
	"ask",
	"go",
	"word",
	"world",
	"question",
	"success",
	"fail",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"test",
	"basketball",
	"stepback",
	"random",
	"characteristic",
	"nevertheless",
	"accommodate",
	"phenomenon",
	"considerable",
	"authropologist",
	"extravaganza",
	"magliano",
	"allege",
	"judicial",
	"likewise",
	"fruitful",
	"legislation",
	"vanguard",
	"situated",
	"rack",
	"contrast",
	"antitrast",
	"blacksmith",
	"lottery",
	"helicopter",
	"ocarina",
	"flamingo",
	"crocodile",
	"panda",
}

func MakeQuiz() []string {
	rand.Seed(time.Now().UnixNano())
	quiz := rand.Perm(len(wordList))
	var quizlist []string
	for _, q := range quiz {
		word := wordList[q]
		quizlist = append(quizlist, word)
	}
	return quizlist
}
