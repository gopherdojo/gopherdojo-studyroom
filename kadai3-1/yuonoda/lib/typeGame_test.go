package typeGame_test

import (
	typeGame "github.com/yuonoda/ggopherdojo-studyroom/kadai3-1/yuonoda/lib/typeGame/lib"
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
	typeGame.Start(10, words)

	cases := []struct {
		name string
		//question []string
		answer []string
		score  int
	}{
		{
			name: "basic",
			//question: []string{"peach", "orange", "apple"},
			//answer: []string{"peach", "orange", "apple"},
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
			},
			score: 10,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

		})
	}
}
