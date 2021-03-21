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
		{name: "jpeg", inputArgs: args{dir: "../testdata/jpeg", oldExt: ".jpeg", newExt: ".png"}},
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

func TestValidateExtParam(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "png", input: ".png", expected: ".png"},
		{name: "uppercase", input: ".PNG", expected: ".png"},
		{name: "no dot", input: "jpeg", expected: ".jpeg"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual, _ := validateExtArg(c.input); actual != c.expected {
				t.Errorf("want validateExtArg(%v) = %v, got %v",
					c.input, c.expected, actual)
			}
		})
	}
}
