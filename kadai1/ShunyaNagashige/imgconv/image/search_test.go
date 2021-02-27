package image_test

import (
	"fmt"
	"testing"

	"github.com/ShunyaNagashige/imgconv/image"
)

type SearchInput struct {
	dir    string
	format string
}

func TestSearch(t *testing.T) {
	cases := []struct {
		name     string
		input    SearchInput
		expected []string
	}{
		{
			name:     "ok",
			input:    SearchInput{dir: ".", format: "png"},
			expected: []string{"testdata/dog_hachi_sasareta.png"},
		},
		//(メモ)search関数内のfilepath.Walkがどうすればエラーを吐き出すかわからない
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Helper()

			actual, _ := image.Search(c.input.dir, c.input.format)

			if len(actual) != len(c.expected) {
				t.Errorf("実際の出力と期待値の,len()の値が異なります: want Search(%s,%s) = %s, got %s",
					c.input.dir, c.input.format, c.expected, actual)
			}

			for i := 0; i < len(c.expected); i++ {
				if actual[i] != c.expected[i] {
					t.Errorf("actual[%d]=%s と c.expected[%d]=%sの値が一致しません: want Search(%s,%s) = %s, got %s",
						i, actual[i], i, c.expected[i], c.input.dir, c.input.format, c.expected, actual)
				}
			}
		})
	}
}

func ExampleSearch(){
	fmt.Println(image.Search(".","png"))
	// output: [testdata/dog_hachi_sasareta.png] <nil>
}
