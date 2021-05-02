package main

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"
)

// requests_test.goで作成しているテストサーバを利用してテストを行う
func TestRun(t *testing.T) {
	ctx := context.Background()

	url := ts.URL
	args := []string{fmt.Sprintf("%s/%s", url, "header.jpg")}
	targetDir := "testdata/test_download"
	timeout := 30 * time.Second

	p := New()
	if err := p.Run(ctx, args, targetDir, timeout); err != nil {
		t.Errorf("failed to Run: %s", err)
	}

	if err := os.Remove(p.FullFileName()); err != nil {
		t.Errorf("failed to remove of result file: %s", err)
	}
}

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
