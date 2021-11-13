package questions

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Vocab struct {
	Words    []string
	Meanings []string
}

// Read words from word file
func ReadWords(filePath string) *Vocab {
	var vc Vocab
	f, _ := os.Open(filePath)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		vc.Words = append(vc.Words, strings.Split(scanner.Text(), ":")[0])
		vc.Meanings = append(vc.Meanings, strings.Split(scanner.Text(), ":")[1])
	}
	return &vc
}

// Get random words from word list
func CreateProblem(vc *Vocab) int {
	// random seed
	rand.Seed(time.Now().UnixNano())
	// get idx of target word
	idx := rand.Intn(len(vc.Words))
	return idx //, vc.words[idx]
}
