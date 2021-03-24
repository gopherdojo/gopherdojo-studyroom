package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/kimuson13/gopherdojo-studyroom/kimuson13/typing/word"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	q := rand.Intn(len(word.WordList))
	fmt.Fprintf(os.Stdout, word.WordList[q])
}
