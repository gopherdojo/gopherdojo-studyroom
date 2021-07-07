package typing

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"testing"
	"time"
)

//go:embed gamedata/words.csv
var words string
var wordsSlice = strings.Split(words, ",")

func TestGame_WithinTheTimelimit(t *testing.T) {
	t.Helper()

	cases := map[string]struct {
		tm       time.Duration
		ans      []string
		expected int
	}{
		"No typo": {
			tm:   1 * time.Second,
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
		"One typo": {
			tm:   1 * time.Second,
			ans: []string{
				"America",
				"American",
				"😊",  // typo
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
		"Two typo" : {
			tm:   3 * time.Second,
			ans: []string{
				"America",
				"American",
				"Angle",
				"April",
				"August",
				"Bacon",
				"🤣",  // tyop 1
				"Battery",
				"👍",  // typo 2
				"Bill",
			},
			expected: 8,
		},
	}

	// []string{"America","American","Angle","April","August","Bacon","Barber","Battery","Bible","Bill"}
	
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			output := new(bytes.Buffer)
			input := bytes.NewBufferString(strings.Join(c.ans, "\n"))
			actual, _ := Game(input, output, wordsSlice, c.tm, true)
			if actual != c.expected {
				t.Errorf("expected %d, but got %d", c.expected, actual)
			}
		})
	}
}

func TestGame_TimelimitOver(t *testing.T) {
	cases := map[string][]string {
		"case 1" : {
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
	}

	for name, val := range cases {
		t.Run(name, func(t *testing.T) {
			output := new(bytes.Buffer)
			input := bytes.NewBufferString(strings.Join(val, "\n"))
			actual, _ := Game(input, output, wordsSlice, time.Duration(1), true)
			if actual > 0 {
				fmt.Println(output.String())
				t.Fatalf("expected 0, but got %d", actual)
			}
			if output.String() != fmt.Sprintf("> %s\n\nGame ends!\nThe number of correct answers is 0\n", val[0]) {
				t.Fatal("message you got is not equal to expected")
			}
		})
	}
}

func Test_input(t *testing.T) {
	t.Helper()

	cases := map[string]string {
		"case 1" : "foo\nbar\nfoobar",
		"case 2" : "foobarfoobar\nbar\nbar",
	}

	for name, val := range cases {
		t.Run(name, func(t *testing.T) {
			inputter := bytes.NewBufferString(val)
			ch := input(inputter)
			expects := strings.Split(val, "\n")

			for _, e := range expects {
				ans := <-ch
				fmt.Println(ans)
				if e != ans {
					t.Fatalf("expected %s, but got %s", e, ans)
				}
			}
		})
	}
}

func Test_incSlice(t *testing.T) {
	t.Helper()

	cases := map[string][]int {
		"case 1 n=5" : {0,1,2,3,4},
		"case 2 n=10" : {0,1,2,3,4,5,6,7,8,9},
	}

	for name, val := range cases {
		t.Run(name, func(t *testing.T) {
			var i int
			for _, v := range val {
				if i != v {
					t.Fatalf("expected %d, but got %d", i, v)
				}
				i++
			}
		})
	}
}
