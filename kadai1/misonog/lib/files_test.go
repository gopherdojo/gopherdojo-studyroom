package lib

import (
	"reflect"
	"testing"
)

func TestExistDir(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "Dir exsit", input: "../testdata/png", expected: true},
		{name: "Dir not exist", input: "../testdata/notexist", expected: false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := ExistDir(c.input); c.expected != actual {
				t.Errorf("want ExitDir(%v) = %v, got %v",
					c.input, c.expected, actual)
			}
		})
	}
}

func TestDirWalk(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{name: "No recursive pattern", input: "../testdata/jpg", expected: []string{"1.jpg", "2.jpg", "3.jpg"}},
		{name: "Recursive pattern", input: "../testdata/png", expected: []string{"1.jpg", "2.jpg", "3.jpg", "recursive/4.png"}},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := dirWalk(c.input); reflect.DeepEqual(c.expected, actual) {
				t.Errorf("want DirWalk(%v) = %v, got %v",
					c.input, c.expected, actual)
			}
		})
	}
}
