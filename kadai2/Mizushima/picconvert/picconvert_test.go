package picconvert_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	picconvert "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai2/Mizushima/picconvert"
)


func TestGlob(t *testing.T) {
	tmpTestDir := "../testdata"

	// make test cases
	cases := []struct {
		name      string
		path      string
		format    []string
		expected1 []string
		expected2 error
	}{
		{name: "jpeg files",
			path:      tmpTestDir,
			format:    []string{"jpg", "jpeg"},
			expected1: []string{tmpTestDir + "/test01.jpg", tmpTestDir + "/" + filepath.Base(tmpTestDir) + "/test01.jpg"},
			expected2: nil,
		},
		{name: "png files",
			path:      tmpTestDir,
			format:    []string{"png"},
			expected1: []string{tmpTestDir + "/test01.png", tmpTestDir + "/" + filepath.Base(tmpTestDir) + "/test01.png"},
			expected2: nil,
		},
		{name: "gif files",
			path:      tmpTestDir,
			format:    []string{"gif"},
			expected1: []string{tmpTestDir + "/test01.gif", tmpTestDir + "/" + filepath.Base(tmpTestDir) + "/test01.gif"},
			expected2: nil,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			actual, err := picconvert.Glob(c.path, c.format)
			// if actual != nil { fmt.Println(actual) }
			if !reflect.DeepEqual(actual, c.expected1) && (err == nil && c.expected2 == nil) {
				t.Errorf("want Glob(%s, %s) = %v, got %v", c.path, c.format, c.expected1, actual)
			}
		})
	}
}

// TestPicConverter_Conv is the test for (p PicConverter)Conv()
func TestPicConverter_Conv(t *testing.T) {
	tmpTestDir := "../testdata"

	// make test cases
	cases := []struct {
		name     string
		pc       *picconvert.PicConverter
		expected error
	}{
		{name: "pre: jpeg, after: png", pc: picconvert.NewPicConverter(tmpTestDir, "jpeg", "png"), expected: nil},
		{name: "pre: jpeg, after: gif", pc: picconvert.NewPicConverter(tmpTestDir, "jpeg", "gif"), expected: nil},
		{name: "pre: jpg, after: png", pc: picconvert.NewPicConverter(tmpTestDir, "jpg", "png"), expected: nil},
		{name: "pre: jpg, after: gif", pc: picconvert.NewPicConverter(tmpTestDir, "jpg", "gif"), expected: nil},
		{name: "pre: png, after: jpg", pc: picconvert.NewPicConverter(tmpTestDir, "png", "jpg"), expected: nil},
		{name: "pre: png, after: gif", pc: picconvert.NewPicConverter(tmpTestDir, "png", "gif"), expected: nil},
		{name: "pre: gif, after: jpg", pc: picconvert.NewPicConverter(tmpTestDir, "gif", "jpg"), expected: nil},
		{name: "pre: gif, after: png", pc: picconvert.NewPicConverter(tmpTestDir, "gif", "png"), expected: nil},
		{name: "pre: jpg, after: xls", pc: picconvert.NewPicConverter(tmpTestDir, "jpg", "xls"), expected: fmt.Errorf("xls is not supported")},
		{name: "no file", pc: picconvert.NewPicConverter(tmpTestDir+"/test02.png", "gif", "png"), expected: fmt.Errorf("there's no [gif] file")},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			actual := c.pc.Conv()
			// if actual != nil { fmt.Println(actual) }
			if actual != nil && actual.Error() != c.expected.Error() {
				t.Errorf("want p.Conv() = %v, got %v", c.expected, actual)
			} else if _, err := os.Stat(fmt.Sprintf("%s/test01_converted.%s", tmpTestDir, c.pc.AfterFormat)); err != nil && c.pc.AfterFormat != "xls" {
				t.Errorf("%s file wasn't made", c.pc.AfterFormat)
			}
		})
	}

	testDeleteConveterd(t, tmpTestDir)
}

// testDeleteConveterd returns a slice of the file paths that meets the "format".
func testDeleteConveterd(t *testing.T, path string) {
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if strings.Contains(path, "_converted") && !info.IsDir() {
				os.Remove(path)
			}
			return nil
		})

	if err != nil {
		t.Error(err)
	}
}

func TestNewPicConverter(t *testing.T) {

	cases := []struct {
		name     string
		path     string
		pre      string
		post     string
		expected *picconvert.PicConverter
	}{
		{
			name:     "pre: jpg, post: png",
			path:     "../testdata",
			pre:      "jpg",
			post:     "png",
			expected: &picconvert.PicConverter{"../testdata", []string{"jpeg", "jpg"}, "png"},
		},
		{
			name:     "pre: jpeg, post: png",
			path:     "../testdata",
			pre:      "jpeg",
			post:     "png",
			expected: &picconvert.PicConverter{"../testdata", []string{"jpeg", "jpg"}, "png"},
		},
		{
			name:     "pre: jpg, post: gif",
			path:     "../testdata",
			pre:      "jpg",
			post:     "gif",
			expected: &picconvert.PicConverter{"../testdata", []string{"jpeg", "jpg"}, "gif"},
		},
		{
			name:     "pre: png, post: jpeg",
			path:     "../testdata",
			pre:      "png",
			post:     "jpeg",
			expected: &picconvert.PicConverter{"../testdata", []string{"png"}, "jpeg"},
		},
		{
			name:     "pre: png, post: jpg",
			path:     "../testdata",
			pre:      "png",
			post:     "jpg",
			expected: &picconvert.PicConverter{"../testdata", []string{"png"}, "jpg"},
		},
		{
			name:     "pre: png, post: gif",
			path:     "../testdata",
			pre:      "png",
			post:     "gif",
			expected: &picconvert.PicConverter{"../testdata", []string{"png"}, "gif"},
		},
		{
			name:     "pre: gif, post: jpeg",
			path:     "../testdata",
			pre:      "gif",
			post:     "jpeg",
			expected: &picconvert.PicConverter{"../testdata", []string{"gif"}, "jpeg"},
		},
		{
			name:     "pre: gif, post: jpg",
			path:     "../testdata",
			pre:      "gif",
			post:     "jpg",
			expected: &picconvert.PicConverter{"../testdata", []string{"gif"}, "jpg"},
		},
		{
			name:     "pre: gif, post: png",
			path:     "../testdata",
			pre:      "gif",
			post:     "png",
			expected: &picconvert.PicConverter{"../testdata", []string{"gif"}, "png"},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			actual := picconvert.NewPicConverter(c.path, c.pre, c.post)
			// if actual != nil { fmt.Println(actual) }
			if !reflect.DeepEqual(actual, c.expected) {
				t.Errorf("want NewPicConverter(%s, %s, %s) = %v, but got %v",
					c.path, c.pre, c.post, c.expected, actual)
			}
		})
	}
}
