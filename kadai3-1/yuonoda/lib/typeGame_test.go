package typeGame_test

import (
	"bytes"
	typeGame "github.com/yuonoda/ggopherdojo-studyroom/kadai3-1/yuonoda/lib/typeGame/lib"
	"strings"
	"testing"
)

var words = []string{
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

func TestStart(t *testing.T) {
	// 設問を定義
	var words = []string{
		"peach",
		"orange",
		"apple",
		"grape",
		"pineapple",
		"mandarin",
		"lemon",
		"kiwi",
		"grapefruit",
		"melon",
	}

	// テストケースを作成
	cases := []struct {
		name   string
		answer []string
		score  int
	}{
		{
			name: "perfect",
			answer: []string{
				"peach",
				"orange",
				"apple",
				"grape",
				"pineapple",
				"mandarin",
				"lemon",
				"kiwi",
				"grapefruit",
				"melon",
			},
			score: 10,
		},
		{
			name: "mistakes",
			answer: []string{
				"peach",
				"orenge",
				"apple",
				"grape",
				"pineapple",
				"mandarin",
				"lemon",
				"kiwi",
				"grapfruit",
				"melon",
			},
			score: 8,
		},
	}

	for _, c := range cases {
		inputLines := strings.Join(c.answer, "\n")
		inputBuf := bytes.NewBufferString(inputLines)
		outputBuf := bytes.NewBuffer([]byte{})
		t.Run(c.name, func(t *testing.T) {
			score := typeGame.Start(10, words, inputBuf, outputBuf)
			if score != c.score {
				t.Error("score doesn't match")
			}
		})
	}
}
