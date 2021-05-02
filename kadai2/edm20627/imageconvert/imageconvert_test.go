package imageconvert_test

import (
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"reflect"
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
		{name: "get jpg", dirs: []string{"../testdata"}, from: "jpg", expected: []string{"../testdata/hoge/img_2.jpg", "../testdata/img_1.jpg"}},
		{name: "get png", dirs: []string{"../testdata"}, from: "png", expected: []string{"../testdata/hoge/img_2.png", "../testdata/img_1.png"}},
		{name: "get gif", dirs: []string{"../testdata"}, from: "gif", expected: []string{"../testdata/hoge/img_2.gif", "../testdata/img_1.gif"}},
	}

	for _, c := range success_cases {
		for _, path := range c.expected {
			err := testCreateImage(t, path, c.from)
			if err != nil {
				t.Fatal("Failed to create the test image")
			}
			defer testRemoveTestdata(t)
		}
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
		name      string
		filepaths []string
		from      string
		to        string
	}{
		{name: "jpg to png", filepaths: []string{"../testdata/img_1.jpg"}, from: "jpg", to: "png"},
		{name: "png to gif", filepaths: []string{"../testdata/img_1.png"}, from: "png", to: "gif"},
		{name: "gif to jpg", filepaths: []string{"../testdata/img_1.gif"}, from: "gif", to: "jpg"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for _, path := range c.filepaths {
				err := testCreateImage(t, path, c.from)
				if err != nil {
					t.Fatal("Failed to create the test image")
				}
				defer testRemoveTestdata(t)
			}
			ci := imageconvert.ConvertImage{Filepaths: c.filepaths, To: c.to}
			if actual := ci.Convert(); actual != nil {
				t.Error(actual)
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
		{name: "jpg to png", from: "jpg", to: "png", expected: true},
		{name: "png to jpg", from: "png", to: "jpg", expected: true},
		{name: "jpeg to gif", from: "jpeg", to: "gif", expected: true},
		{name: "gif to jpeg", from: "gif", to: "jpeg", expected: true},
		{name: "hoge to fuga", from: "hoge", to: "fuga", expected: false},
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

func testCreateImage(t *testing.T, path string, fmt string) error {
	t.Helper()

	err := os.MkdirAll("../testdata/hoge", 0777)
	if err != nil {
		t.Error(err)
	}

	file, err := os.Create(path)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for i := img.Rect.Min.Y; i < img.Rect.Max.Y; i++ {
		for j := img.Rect.Min.X; j < img.Rect.Max.X; j++ {
			img.Set(j, i, color.RGBA{255, 255, 0, 0})
		}
	}

	switch fmt {
	case "png":
		if err := png.Encode(file, img); err != nil {
			return err
		}
	case "jpg", "jpeg":
		if err := jpeg.Encode(file, img, nil); err != nil {
			return err
		}
	case "gif":
		if err := gif.Encode(file, img, nil); err != nil {
			return err
		}
	}

	return nil
}

func testRemoveTestdata(t *testing.T) error {
	t.Helper()

	err := os.RemoveAll("../testdata")
	if err != nil {
		t.Error(err)
	}

	return nil
}
