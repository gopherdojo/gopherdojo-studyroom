package conversion

import "testing"

type args struct {
	diraName     string
	outDirectory string
	beforeExt    string
	afterExt     string
}

func TestConvert(t *testing.T) {
	cases := []struct {
		name  string
		input args
	}{
		{name: "convert png from jpg", input: args{diraName: "../image", outDirectory: "../output", beforeExt: "jpg", afterExt: "png"}},
		{name: "convert jpg from png", input: args{diraName: "../image", outDirectory: "../output", beforeExt: "png", afterExt: "jpg"}},
		{name: "convert gif from jpg", input: args{diraName: "../image", outDirectory: "../output", beforeExt: "jpg", afterExt: "gif"}},
		{name: "convert gif from png", input: args{diraName: "../image", outDirectory: "../output", beforeExt: "png", afterExt: "gif"}},
		{name: "convert hoge from png", input: args{diraName: "../image", outDirectory: "../output", beforeExt: "png", afterExt: "hoge"}},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := Convert(c.input.diraName, c.input.outDirectory, c.input.beforeExt, c.input.afterExt); actual != nil {
				t.Errorf("Received unexpected error: \n%v", actual)
			}
		})
	}
}

func TestGetFileNameWithoutExt(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		output string
	}{
		{name: "get filename without ext", input: "./image/sample1.jpg", output: "sample"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := getFileNameWithoutExt(c.input); actual == c.output {
				t.Errorf("Received unexpected error: \n%v", actual)
			}
		})
	}
}
