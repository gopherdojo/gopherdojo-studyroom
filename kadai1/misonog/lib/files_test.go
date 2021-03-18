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
		{name: "No recursive pattern", input: "../testdata/jpg", expected: []string{"../testdata/jpg/1.jpg", "../testdata/jpg/2.jpg", "../testdata/jpg/3.jpg"}},
		{name: "Recursive pattern", input: "../testdata/png", expected: []string{"../testdata/png/1.png", "../testdata/png/2.png", "../testdata/png/3.png", "../testdata/png/recursive/4.png"}},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := dirWalk(c.input); !reflect.DeepEqual(c.expected, actual) {
				t.Errorf("want DirWalk(%v) = %v, got %v",
					c.input, c.expected, actual)
			}
		})
	}
}

func TestGetFileStruct(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		expected []File
	}{
		{name: "Single file pattern", input: []string{"1.jpg"}, expected: []File{{Path: "1.jpg", Ext: ".jpg"}}},
		{name: "Multi file pattern", input: []string{"1.jpg", "2.png"}, expected: []File{{Path: "1.jpg", Ext: ".jpg"}, {Path: "2.png", Ext: ".png"}}},
		{name: "Blank pattern", input: []string{""}, expected: []File{{Path: "", Ext: ""}}},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := getFileStruct(c.input); !reflect.DeepEqual(c.expected, actual) {
				t.Errorf("want getFileStruct(%v) = %v, got %v",
					c.input, c.expected, actual)
			}
		})
	}
}
