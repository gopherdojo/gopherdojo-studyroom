package utilities_test

import (
	"github.com/yuonoda/gopherdojo-studyroom/kadai3-2/yuonoda/utilities"
	"reflect"
	"testing"
)

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
			utilities.FillByteArr(c.arr[:], c.startAt, c.partArr)
			if !reflect.DeepEqual(c.expectedArr, c.arr) {
				t.Error("Array does not match")
			}
		})
	}

}
