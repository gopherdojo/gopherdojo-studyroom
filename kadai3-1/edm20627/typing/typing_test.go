package typing_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/edm20627/gopherdojo-studyroom/kadai3-1/edm20627/typing"
)

var words = []string{
	"go",
	"java",
	"ruby",
	"php",
	"javascript",
	"python",
	"kotlin",
	"swift",
	"c",
}

func TestStart(t *testing.T) {
	cases := []struct {
		name     string
		gameTime time.Duration
		answer   []string
		score    int
	}{
		{
			name:     "success",
			gameTime: 20 * time.Second,
			answer: []string{
				"go",
				"java",
				"ruby",
				"php",
				"javascript",
				"python",
				"kotlin",
				"swift",
				"c",
			},
			score: 9,
		},
		{
			name:     "2 typos",
			gameTime: 20 * time.Second,
			answer: []string{
				"typo1",
				"typo2",
				"ruby",
				"php",
				"javascript",
				"python",
				"kotlin",
				"swift",
				"c",
			},
			score: 7,
		},
		{
			name:     "timeout",
			gameTime: 0 * time.Second,
			answer: []string{
				"go",
				"java",
				"ruby",
				"php",
				"javascript",
				"python",
				"kotlin",
				"swift",
				"c",
			},
			score: 0,
		},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {
			input := bytes.NewBufferString(strings.Join(c.answer, "\n"))
			output := new(bytes.Buffer)
			score := typing.Start(input, output, words, c.gameTime)
			if c.score != score {
				t.Errorf("expected %d, but got %d", c.score, score)
			}
		})
	}
}
