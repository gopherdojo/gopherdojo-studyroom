package imgconv

import (
	"testing"
)

func TestGetConvertImages(t *testing.T) {
	dir := "../testdata"
	cases := []struct {
		from, to string
		expected []ConvertImage
	}{
		{
			from: "jpg",
			to:   "png",
			expected: []ConvertImage{
				{FromPath: "../testdata/A.jpg", ToPath: "../testdata/A.png"},
				{FromPath: "../testdata/B.jpg", ToPath: "../testdata/B.png"},
			},
		},
		{
			from: "png",
			to:   "jpg",
			expected: []ConvertImage{
				{FromPath: "../testdata/C.png", ToPath: "../testdata/C.jpg"},
				{FromPath: "../testdata/D.png", ToPath: "../testdata/D.jpg"},
			},
		},
		{
			from: "gif",
			to:   "png",
			expected: []ConvertImage{
				{FromPath: "../testdata/sub/E.gif", ToPath: "../testdata/sub/E.png"},
				{FromPath: "../testdata/sub/F.gif", ToPath: "../testdata/sub/F.png"},
			},
		},
	}

	for _, c := range cases {
		images, _ := GetConvertImages(dir, c.from, c.to)
		for index, image := range images {
			if image.FromPath != c.expected[index].FromPath {
				t.Errorf("FromPath=%s, want %s", image.FromPath, c.expected[index].FromPath)
			}
			if image.ToPath != c.expected[index].ToPath {
				t.Errorf("ToPath=%s, want %s", image.ToPath, c.expected[index].ToPath)
			}
		}
	}
}
