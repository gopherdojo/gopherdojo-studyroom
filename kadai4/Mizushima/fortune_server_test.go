package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestOmikujiHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	omikujiHandler(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	fmt.Println(string(b))
}

func TestResult(t *testing.T) {
	cases := map[string]struct {
		date     time.Time
		index    int
		expected string
	}{
		"the 1st of the first three days of the new year.": {
			date:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
			index:    1,
			expected: "大吉",
		},
		"the 2nd of the first three days of the new year.": {
			date:     time.Date(2021, 1, 2, 0, 0, 0, 0, time.Local),
			index:    1,
			expected: "大吉",
		},
		"the 3rd of the first three days of the new year.": {
			date:     time.Date(2021, 1, 3, 0, 0, 0, 0, time.Local),
			index:    1,
			expected: "大吉",
		},
		"the 4th of the first three days of the new year, and the index is 5.": {
			date:     time.Date(2021, 1, 4, 0, 0, 0, 0, time.Local),
			index:    5,
			expected: "凶",
		},
		"August 25 and the index is 3": {
			date:     time.Date(2021, 8, 25, 0, 0, 0, 0, time.Local),
			index:    3,
			expected: "小吉",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			actual := result(c.index, c.date)
			if actual != c.expected {
				t.Fatalf("expected %v, but got %v", c.expected, actual)
			}
		})
	}
}
