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
		{name: "+odd", input: args{diraName: "./image", outDirectory: "./output", beforeExt: "jpg", afterExt: "png"}},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if actual := Convert(c.input.diraName, c.input.outDirectory, c.input.beforeExt, c.input.afterExt); actual != nil {
				t.Errorf("Received unexpected error:\n%v", actual)
			}
		})
	}
}
