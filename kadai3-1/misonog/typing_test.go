package main

import (
	"testing"
)

func TestGetPoint(t *testing.T) {
	cases := []struct {
		name     string
		expected int
	}{
		{name: "get", expected: 0},
	}

	for _, c := range cases {
		typing := Typing{}

		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := typing.getPoint(); c.expected != actual {
				t.Errorf("want typing.getPoint() = %v, got %v",
					c.expected, actual)
			}
		})
	}
}

func TestCheck(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "true", input: "Go", expected: true},
		{name: "blank", input: "", expected: false},
	}

	for _, c := range cases {
		typing := Typing{Word: "Go"}

		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := typing.check(c.input); c.expected != actual {
				t.Errorf("want typing.check(%v) = %v, got %v",
					c.input, c.expected, actual)
			}
		})
	}
}
