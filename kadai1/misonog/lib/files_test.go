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
		{name: "dir exsit", input: "../testdata/png", expected: true},
		{name: "dir not exist", input: "../testdata/notexist", expected: false},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := existDir(c.input); c.expected != actual {
				t.Errorf("want exitDir(%v) = %v, got %v",
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
		{name: "no recursive", input: "../testdata/jpg", expected: []string{"../testdata/jpg/1.jpg", "../testdata/jpg/2.jpg", "../testdata/jpg/3.jpg"}},
		{name: "recursive", input: "../testdata/png", expected: []string{"../testdata/png/1.png", "../testdata/png/2.png", "../testdata/png/3.png", "../testdata/png/recursive/4.png"}},
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
		{name: "single file", input: []string{"1.jpg"}, expected: []File{{Path: "1.jpg", Ext: ".jpg"}}},
		{name: "multi file", input: []string{"1.jpg", "2.png"}, expected: []File{{Path: "1.jpg", Ext: ".jpg"}, {Path: "2.png", Ext: ".png"}}},
		{name: "blank", input: []string{""}, expected: []File{{Path: "", Ext: ""}}},
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

func TestFilter(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected int
	}{
		{name: "filter", input: ".png", expected: 1},
		{name: "filter jpg", input: ".jpg", expected: 1},
		{name: "no filter", input: ".gif", expected: 0},
		{name: "blank", input: "blank", expected: 0},
	}

	for _, c := range cases {
		var fileList Files
		f := []File{{Path: "1.png", Ext: ".png"}, {Path: "2.jpg", Ext: ".jpg"}}
		for _, file := range f {
			fileList = append(fileList, file)
		}
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Run(c.name, func(t *testing.T) {
				if actual := fileList.filter(c.input); c.expected != len(actual) {
					t.Errorf("want getFileStruct(%v) = %v, got %v",
						c.input, c.expected, actual)
				}
			})
		})
	}
}

func TestCmpExt(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "match ext", input: ".png", expected: true},
		{name: "unmatch ext", input: ".jpg", expected: false},
		{name: "blank ext", input: "", expected: false},
	}

	for _, c := range cases {
		f := File{Path: "test.png", Ext: ".png"}
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := f.cmpExt(c.input); c.expected != actual {
				t.Errorf("want File.cmpExt(%v) = %v, got %v",
					c.input, c.expected, actual)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	// TODO: テストの書き方がわからないので、一旦エラーを吐かないかでお茶を濁す
	cases := []struct {
		name  string
		input string
		f     File
	}{
		{name: "png", input: "png", f: File{Path: "../testdata/jpeg/1.jpeg", Ext: "jpeg"}},
		{name: "gif", input: "gif", f: File{Path: "../testdata/jpeg/1.jpeg", Ext: "jpeg"}},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := c.f.convert(c.input); actual != nil {
				t.Errorf("Received unexpected error:\n%v", actual)
			}
		})
	}
}
