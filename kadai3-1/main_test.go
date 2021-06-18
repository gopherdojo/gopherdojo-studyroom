package main

import (
	"testing"
)

func TestRandomWord(t *testing.T) {
	word := RandomWord()
	if word == nil {
		t.Error("文字取り出さない")
	}
}
