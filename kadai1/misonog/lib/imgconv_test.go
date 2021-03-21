package lib

import "testing"

type args struct {
	dir    string
	oldExt string
	newExt string
}

func TestImgConv(t *testing.T) {
	cases := []struct {
		name      string
		inputArgs args
	}{
		{name: "png", inputArgs: args{dir: "../testdata/png/recursive", oldExt: ".png", newExt: ".jpg"}},
		{name: "png recursive", inputArgs: args{dir: "../testdata/png", oldExt: ".png", newExt: ".jpg"}},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			actual := ImgConv(c.inputArgs.dir, c.inputArgs.oldExt, c.inputArgs.newExt)
			if actual != nil {
				t.Errorf("Received unexpected error:\n%v", actual)
			}
		})
	}
}
