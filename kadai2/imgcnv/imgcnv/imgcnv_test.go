package imgcnv_test

import (
	"errors"
	"testing"

	"github.com/dai65527/gopherdojo-studyroom/kadai1/imgcnv"
)

func checkError(t *testing.T, expected, actual error) {
	if actual != expected && actual.Error() != expected.Error() {
		t.Errorf("actual: %#v, expected: %#v", actual.Error(), expected.Error())
	}
}

func TestConvert(t *testing.T) {
	cases := []struct {
		name     string
		filename string
		in, out  imgcnv.Extension
		expected error
	}{
		{name: "png to jpg", filename: "../images/image.png", in: "png", out: "jpg", expected: nil},
		{name: "jpg to png", filename: "../images/image.jpg", in: "jpg", out: "png", expected: nil},
		{name: "jpeg to png", filename: "../images/image.jpeg", in: "jpg", out: "png", expected: nil},
		{name: "invalid input extension", filename: "../images/image.hoge", in: "hoge", out: "png", expected: errors.New("invalid input file extension")},
		{name: "invalid output extension", filename: "../images/image.jpg", in: "jpg", out: "hoge", expected: errors.New("invalid output file extension")},
		{name: "invalid file name", filename: "nonexisitfile.jpg", in: "jpg", out: "png", expected: errors.New("failed to open file")},
		{name: "invalid jpg file", filename: "imgcnv_test.go", in: "jpg", out: "png", expected: errors.New("invalid JPEG format: missing SOI marker")},
		{name: "invalid png file", filename: "imgcnv_test.go", in: "png", out: "jpg", expected: errors.New("png: invalid format: not a PNG file")},
	}

	for _, c := range cases {
		checkError(t, c.expected, imgcnv.Convert(c.filename, c.in, c.out))
	}
}
