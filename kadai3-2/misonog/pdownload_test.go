package main

import (
	"testing"
)

func TestParseURL(t *testing.T) {

	cases := []struct {
		name     string
		input    []string
		expected string
	}{
		{name: "an URL", input: []string{"https://www.google.com/"}, expected: "https://www.google.com/"},
		{name: "URLs", input: []string{"https://www.google.com/", "https://golang.org/"}, expected: ""},
		{name: "invalid URL", input: []string{"invalid_url"}, expected: ""},
	}

	for _, c := range cases {
		c := c
		p := New()
		t.Run(c.name, func(t *testing.T) {
			_ = p.parseURL(c.input)
			if actual := p.URL; c.expected != actual {
				t.Errorf("want p.URL = %v, got %v",
					c.expected, actual)
			}
		})
	}
}
