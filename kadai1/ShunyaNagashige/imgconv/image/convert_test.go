package image_test

import (
	"testing"

	"github.com/ShunyaNagashige/imgconv/image"
)

type ConvertInput struct {
	src       string
	srcFormat string
	dstFormat string
}

func TestConvert(t *testing.T) {
	cases := []struct {
		name     string
		input    ConvertInput
		expected error
	}{
		{
			name:     "png:jpg",
			input:    ConvertInput{src: "testdata/dog_hachi_sasareta.png", srcFormat: "png", dstFormat: "jpg"},
			expected: nil,
		},
		{
			name:     "png:gif",
			input:    ConvertInput{src: "testdata/dog_hachi_sasareta.png", srcFormat: "png", dstFormat: "gif"},
			expected: nil,
		},
		{
			name:     "jpg:png",
			input:    ConvertInput{src: "testdata/itu.jpg", srcFormat: "jpg", dstFormat: "png"},
			expected: nil,
		},
		{
			name:     "jpg:gif",
			input:    ConvertInput{src: "testdata/itu.jpg", srcFormat: "jpg", dstFormat: "gif"},
			expected: nil,
		},
		{
			name:     "gif:png",
			input:    ConvertInput{src: "testdata/pop.gif", srcFormat: "gif", dstFormat: "png"},
			expected: nil,
		},
		{
			name:     "gif:jpg",
			input:    ConvertInput{src: "testdata/pop.gif", srcFormat: "gif", dstFormat: "jpg"},
			expected: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Helper()

			if actual := image.Convert(c.input.src, c.input.srcFormat, c.input.dstFormat); actual != nil {
				t.Error("nil")
			}
		})
	}
}