package splitDownload_test

import (
	"reflect"
	"testing"

	splitDownload "github.com/yuonoda/gopherdojo-studyroom/kadai3-2/yuonoda/lib"
)

func TestGetPartialContent(t *testing.T) {
	url := "https://dumps.wikimedia.org/jawiki/20210101/jawiki-20210101-pages-articles-multistream-index.txt.bz2"
	pcch := make(chan splitDownload.ExportedPartialContent, 10)
	splitDownload.ExportedGetPartialContent(url, 10, 20, pcch)
}

func TestFillByteArr(t *testing.T) {
	cases := []struct {
		name        string
		arr         []byte
		startAt     int
		partArr     []byte
		expectedArr []byte
	}{
		{
			name:        "basic",
			arr:         []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			startAt:     3,
			partArr:     []byte{4, 5, 6},
			expectedArr: []byte{0, 0, 0, 4, 5, 6, 0, 0, 0, 0},
		},
		{
			name:        "basic2",
			arr:         []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			startAt:     8,
			partArr:     []byte{9, 10},
			expectedArr: []byte{0, 0, 0, 0, 0, 0, 0, 0, 9, 10},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			splitDownload.ExportedFillByteArr(c.arr[:], c.startAt, c.partArr)
			if !reflect.DeepEqual(c.expectedArr, c.arr) {
				t.Error("Array does not match")
			}
		})
	}

}
