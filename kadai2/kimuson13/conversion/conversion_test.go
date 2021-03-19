package conversion_test

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"testing"

	"github.com/kimuson13/gopherdojo-studyroom/kimuson13/conversion"
)

var cs conversion.ConvertStruct

func TestExtensionCheck(t *testing.T) {
	cases := []struct {
		name, input string
		expected    error
	}{
		{name: "jpeg", input: "jpeg", expected: nil},
		{name: "jpg", input: "jpg", expected: nil},
		{name: "png", input: "png", expected: nil},
		{name: "gif", input: "gif", expected: nil},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			testExtensionCheck(t, c.input, c.expected)
		})
	}
}

func testExtensionCheck(t *testing.T, input string, expected error) {
	t.Helper()
	err := conversion.ExtensionCheck(input)
	if err != nil {
		t.Fatal(err)
	}
	if err != expected {
		t.Errorf(
			"want ExtensionCheck(%v) = %v, got %v",
			input, expected, err,
		)
	}
}

func TestWalkDirs(t *testing.T) {
	var (
		dirs     []string
		expected error
	)
	if err := os.Mkdir("testdir", 0777); err != nil {
		t.Fatal(err)
	}
	dirs = append(dirs, "testdir")
	if actual := cs.WalkDirs(dirs); actual != expected {
		t.Errorf("walk error")
	}
	if err := os.Remove("testdir"); err != nil {
		t.Fatal(err)
	}
}

func TestConvert(t *testing.T) {
	cases := []struct {
		name, before, after string
		imgId               int
		expected            error
	}{
		{name: "jpegTopng", before: "jpeg", after: "png", imgId: 1, expected: nil},
		{name: "jpegTogif", before: "jpeg", after: "gif", imgId: 2, expected: nil},
		{name: "jpgTopng", before: "jpg", after: "png", imgId: 3, expected: nil},
		{name: "jpgTogif", before: "jpg", after: "gif", imgId: 4, expected: nil},
		{name: "pngTojpeg", before: "png", after: "jpeg", imgId: 5, expected: nil},
		{name: "pngTojpg", before: "png", after: "jpg", imgId: 6, expected: nil},
		{name: "pngTogif", before: "png", after: "gif", imgId: 7, expected: nil},
		{name: "gifTojpeg", before: "gif", after: "jpeg", imgId: 8, expected: nil},
		{name: "gifTojpg", before: "gif", after: "jpg", imgId: 9, expected: nil},
		{name: "gifTopng", before: "gif", after: "png", imgId: 10, expected: nil},
	}
	var dirs []string
	if err := os.Mkdir("testdir2", 0777); err != nil {
		t.Fatal(err)
	}
	dirs = append(dirs, "testdir2")
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			path := testconvert(t, dirs, c.before, c.after, c.imgId, c.expected)
			if err := os.Remove(path); err != nil {
				t.Fatal(err)
			}
		})
	}
	if err := os.RemoveAll("testdir2"); err != nil {
		t.Fatal(err)
	}
}

func testconvert(t *testing.T, dirs []string, before, after string, imageId int, expected error) string {
	t.Helper()
	var id string
	id = strconv.Itoa(imageId)
	img := image.NewNRGBA(image.Rect(255, 255, 0, 0))
	out, err := os.Create("testdir2/image" + id + "." + before)
	defer out.Close()
	if err != nil {
		t.Fatal(err)
	}
	switch before {
	case "jpeg", "jpg":
		if err := jpeg.Encode(out, img, nil); err != nil {
			t.Fatal(err)
		}
	case "png":
		if err := png.Encode(out, img); err != nil {
			t.Fatal(err)
		}
	case "gif":
		if err := gif.Encode(out, img, nil); err != nil {
			t.Fatal(err)
		}
	}
	cs.After = after
	actual := cs.WalkDirs(dirs)
	if actual != nil {
		t.Errorf("walk error")
	}
	actual = cs.Convert()
	if actual != expected {
		t.Fatal(err)
	}
	cs.After = ""
	pathname := "testdir2/image" + id + "." + after
	return pathname
}
