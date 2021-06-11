package main

import (
	"flag"
	"fmt"
	"testing"
)

func TestIsSupportedFormat(t *testing.T) {
	cases := []struct {
		name     string
		format   string
		expected bool
	}{
		{name: "jpg", format: "jpg", expected: true},
		{name: "jpeg", format: "jpeg", expected: true},
		{name: "png", format: "png", expected: true},
		{name: "gif", format: "gif", expected: true},
		{name: "xls", format: "xls", expected: false},
	}

	for _, c := range cases {
		// c := c
		t.Run(c.name, func(t *testing.T) {
			actual := IsSupportedFormat(c.format)
			// if actual != nil { fmt.Println(actual) }
			if actual != c.expected {
				t.Errorf("want IsSupportedFormat(%s) = %v, got %v", c.format, c.expected, actual)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	cases := []struct {
		name     string
		pre      string
		post     string
		path     string
		expected error
	}{
		{name: "jpg and png", pre: "jpg", post: "png", path: ".", expected: nil},
		{name: "jpeg and png", pre: "jpeg", post: "png", path: ".", expected: nil},
		{name: "jpg and gif", pre: "jpg", post: "gif", path: ".", expected: nil},
		{name: "jpeg and gif", pre: "jpeg", post: "gif", path: ".", expected: nil},
		{name: "png and jpg", pre: "png", post: "jpg", path: ".", expected: nil},
		{name: "png and jpeg", pre: "png", post: "jpeg", path: ".", expected: nil},
		{name: "png and gif", pre: "png", post: "gif", path: ".", expected: nil},
		{name: "gif and jpeg", pre: "gif", post: "jpeg", path: ".", expected: nil},
		{name: "gif and jpg", pre: "gif", post: "jpg", path: ".", expected: nil},
		{name: "gif and png", pre: "gif", post: "png", path: ".", expected: nil},
		{name: "xls and gif", pre: "xls", post: "gif", path: ".", expected: fmt.Errorf("-pre xls is not supported")},
		{name: "gif and xls", pre: "gif", post: "xls", path: ".", expected: fmt.Errorf("-post xls is not supported")},
		{
			name:     "jpg and jpeg",
			pre:      "jpg",
			post:     "jpeg",
			path:     ".",
			expected: fmt.Errorf("the parameter of -pre is same as that of -post. 'jpg' is considered same as 'jpeg'"),
		},
		{
			name:     "jpeg and jpg",
			pre:      "jpeg",
			post:     "jpg",
			path:     ".",
			expected: fmt.Errorf("the parameter of -pre is same as that of -post. 'jpg' is considered same as 'jpeg'"),
		},
		{
			name:     "png and png",
			pre:      "png",
			post:     "png",
			path:     ".",
			expected: fmt.Errorf("the parameter of -pre is same as that of -post"),
		},
		{
			name:     "gif and gif",
			pre:      "gif",
			post:     "gif",
			path:     ".",
			expected: fmt.Errorf("the parameter of -pre is same as that of -post"),
		},
	}

	for _, c := range cases {
		// c := c
		t.Run(c.name, func(t *testing.T) {
			flag.CommandLine.Set("pre", c.pre)
			flag.CommandLine.Set("post", c.post)
			flag.CommandLine.Set("path", c.path)
			actual := Validate()
			// if actual != nil { fmt.Println(actual) }
			if actual != nil && c.expected != nil && actual.Error() != c.expected.Error() {
				t.Errorf("want Validate() = %v, but got %v", c.expected.Error(), actual.Error())
			} else if actual != nil && c.expected == nil {
				t.Errorf("want Validate() = nil, but got %v", actual.Error())
			} else if actual == nil && c.expected != nil {
				t.Errorf("want Validate() = %v, but got nil", c.expected.Error())
			}
		})
	}
}
