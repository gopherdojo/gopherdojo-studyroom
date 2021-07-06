package typing_test

import (
	"bytes"
	_ "embed"
	"strings"
	"testing"
	"time"

	typing "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-1/Mizushima/typing"
)

//go:embed gamedata/words.csv
var words string
var wordsSlice = strings.Split(words, ",")

func TestGame(t *testing.T) {
	t.Helper()

	cases := []struct {
		name     string
		tm       time.Duration
		ans      []string
		expected int
	}{
		{
			name: "No typo",
			tm:   3 * time.Second,
			ans: []string{
				"America",
				"American",
				"Angle",
				"April",
				"August",
				"Bacon",
				"Barber",
				"Battery",
				"Bible",
				"Bill",
			},
			expected: 10,
		},
		{
			name: "One typo",
			tm:   3 * time.Second,
			ans: []string{
				"America",
				"American",
				"😊",
				"April",
				"August",
				"Bacon",
				"Barber",
				"Battery",
				"Bible",
				"Bill",
			},
			expected: 9,
		},
	}

	// []string{"America","American","Angle","April","August","Bacon","Barber","Battery","Bible","Bill"}
	
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			output := new(bytes.Buffer)
			input := bytes.NewBufferString(strings.Join(c.ans, "\n"))
			actual, _ := typing.Game(input, output, wordsSlice, c.tm, true)
			if actual != c.expected {
				t.Errorf("wanted %d, but got %d", c.expected, actual)
			}
		})
	}
}
