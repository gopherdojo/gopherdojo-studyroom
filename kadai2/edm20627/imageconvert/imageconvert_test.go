package imageconvert_test

import (
	"errors"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/edm20627/gopherdojo-studyroom/kadai2/edm20627/imageconvert"
)

func TestGet(t *testing.T) {
	success_cases := []struct {
		name      string
		dirs      []string
		filepaths []string
		from      string
		expected  []string
	}{
		{name: "get jpg", dirs: []string{"../testdata"}, from: "jpg", expected: []string{"../testdata/hoge/img_1.jpg", "../testdata/img_1.jpg"}},
		{name: "get png", dirs: []string{"../testdata"}, from: "png", expected: []string{"../testdata/hoge/img_2.png", "../testdata/img_2.png"}},
		{name: "get gif", dirs: []string{"../testdata"}, from: "gif", expected: []string{"../testdata/hoge/img_3.gif", "../testdata/img_3.gif"}},
	}

	for _, c := range success_cases {
		t.Run(c.name, func(t *testing.T) {
			ci := imageconvert.ConvertImage{From: c.from}
			if err := ci.Get(c.dirs); err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(ci.Filepaths, c.expected) {
				t.Errorf("want %v, got %v", c.expected, ci.Filepaths)
			}
		})
	}

	failure_cases := []struct {
		name      string
		dirs      []string
		filepaths []string
		from      string
		expected  error
	}{
		{name: "occurred ErrNotSpecified", dirs: []string{}, from: "jpg", expected: imageconvert.ErrNotSpecified},
		{name: "occurred ErrNotDirectory", dirs: []string{"../non_testdata"}, from: "jpg", expected: imageconvert.ErrNotDirectory},
	}

	for _, c := range failure_cases {
		t.Run(c.name, func(t *testing.T) {
			ci := imageconvert.ConvertImage{From: c.from}
			if err := ci.Get(c.dirs); err != c.expected {
				t.Errorf("want %v, got %v", c.expected, err)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	cases := []struct {
		name         string
		filepaths    []string
		from         string
		to           string
		deleteOption bool
	}{
		{name: "jpg to png", filepaths: []string{"../testdata/img_1.jpg"}, from: "jpg", to: "png", deleteOption: false},
		{name: "png to gif", filepaths: []string{"../testdata/img_2.png"}, from: "png", to: "gif", deleteOption: false},
		{name: "gif to jpg", filepaths: []string{"../testdata/img_3.gif"}, from: "gif", to: "jpg", deleteOption: false},
		{name: "jpg to png", filepaths: []string{"../testdata/img_1.jpg"}, from: "jpg", to: "png", deleteOption: true},
		{name: "png to gif", filepaths: []string{"../testdata/img_2.png"}, from: "png", to: "gif", deleteOption: true},
		{name: "gif to jpg", filepaths: []string{"../testdata/img_3.gif"}, from: "gif", to: "jpg", deleteOption: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			imageconvert.OsRemove = func(path string) error {
				for _, filepath := range c.filepaths {
					if path == filepath {
						return nil
					}
				}
				return errors.New("failed to delete for conversion source image")
			}

			ci := imageconvert.ConvertImage{Filepaths: c.filepaths, To: c.to, DeleteOption: c.deleteOption}
			if actual := ci.Convert(); actual != nil {
				t.Error(actual)
			}
			for _, filepath := range c.filepaths {
				targetFile := strings.Replace(filepath, c.from, c.to, 1)
				if err := os.Remove(targetFile); err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func TestValid(t *testing.T) {
	cases := []struct {
		name     string
		from     string
		to       string
		expected bool
	}{
		{name: "jpg to png (true to true)", from: "jpg", to: "png", expected: true},
		{name: "png to jpg (true to true)", from: "png", to: "jpg", expected: true},
		{name: "jpeg to gif (true to true)", from: "jpeg", to: "gif", expected: true},
		{name: "gif to jpeg (true to true)", from: "gif", to: "jpeg", expected: true},
		{name: "jpg to fuga (true to false)", from: "jpg", to: "fuga", expected: false},
		{name: "hoge to jpg (false to true)", from: "hoge", to: "jpg", expected: false},
		{name: "hoge to fuga (false to false)", from: "hoge", to: "fuga", expected: false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ci := imageconvert.ConvertImage{From: c.from, To: c.to}
			if actual := ci.Valid(); c.expected != actual {
				t.Errorf("want %v, got %v", c.expected, actual)
			}
		})
	}
}
