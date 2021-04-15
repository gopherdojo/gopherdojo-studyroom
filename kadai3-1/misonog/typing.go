package main

import "math/rand"

type Typing struct {
	Words []string
	Word  string
	Point int
}

func (t *Typing) plus() {
	t.Point += 1
}

func (t *Typing) getPoint() int {
	return t.Point
}

func (t *Typing) check(input string) bool {
	if t.Word == input {
		return true
	} else {
		return false
	}
}

func (t *Typing) shuffle() {
	i := rand.Intn(len(t.Words))
	t.Word = t.Words[i]
}
