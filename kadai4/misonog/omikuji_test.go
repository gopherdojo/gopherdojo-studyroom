package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	cases := []struct {
		name        string
		date        string
		fixedResult bool
		expected    string
	}{
		{name: "no date", date: ""},
		{name: "12/31", date: "2020-12-31"},
		{name: "shogatsu", date: "2021-01-01", fixedResult: true, expected: "大吉"},
		{name: "1/4", date: "2021-01-04"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/?date="+c.date, nil)
			handler(w, r)
			rw := w.Result()
			defer rw.Body.Close()
			if rw.StatusCode != http.StatusOK {
				t.Fatal("unexpected status code")
			}

			var res omikujiResult
			dec := json.NewDecoder(rw.Body)
			if err := dec.Decode(&res); err != nil {
				t.Fatal(err)
			}
			if c.fixedResult && c.expected != res.Result {
				t.Errorf("want omikujiResult.Result = %v, got %v",
					res.Result, c.expected)
			}
		})
	}
}

func TestIsShogatsu(t *testing.T) {
	cases := []struct {
		name     string
		input    time.Time
		expected bool
	}{
		{name: "nil", expected: false},
		{name: "shogatsu 1/1", input: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), expected: true},
		{name: "shogatsu 1/2", input: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), expected: true},
		{name: "not shogatsu", input: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC), expected: false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := isShogatsu(c.input); actual != c.expected {
				t.Errorf("want isShogatsu(%v) = %v, got %v",
					c.input, c.expected, actual)
			}
		})
	}
}
