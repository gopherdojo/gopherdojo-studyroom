package imgconv_test

import (
	"imgconv/imgconv"
	"testing"
)

func TestDo (t *testing.T) {
	cases := []struct{
		name string
		filePath string
		ext string
		output error
	}{
		{name : "jpeg to png" , filePath: "../testdata/dst/jpeg/1.jpeg", ext: "png",  output : nil},
		{name : "png to jpeg" , filePath: "../testdata/dst/png/1.png", ext: "jpeg",  output : nil},
	}

	for _, c := range cases {
		conv := &imgconv.ImgConv{FilePath: c.filePath, Ext: c.ext}
		err := conv.Do()
		if err != c.output {
			t.Errorf("invalid result. testCase:%#v, actual:%v", c.output, err)
		}
	}
}
