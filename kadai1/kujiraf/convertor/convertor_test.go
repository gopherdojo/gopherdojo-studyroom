package imgconv

import (
	"testing"
)

var validateTests = []struct {
	name string
	in   Convertor
	out  string
}{
	{
		"no -in",
		Convertor{},
		"-in is required",
	},
	{
		"invalid -from",
		Convertor{
			Src:  "aaa",
			From: "jjpeg",
		},
		"-from dose not support .jjpeg",
	},
	{
		"invalid -to",
		Convertor{
			Src:  "aaa",
			From: "jpeg",
			To:   "ppng",
		},
		"-to dose not support .ppng",
	},
	{
		"-from and -to are same",
		Convertor{
			Src:  "aaa",
			From: "png",
			To:   "png",
		},
		"-from and -to are same. -from .png, -to .png",
	},
	{
		"both -from and -to are jpg",
		Convertor{
			Src:  "aaa",
			From: "jpeg",
			To:   "jpg",
		},
		"-from and -to are same. -from .jpeg, -to .jpg",
	},
}

func TestValidateNG(t *testing.T) {
	for _, tt := range validateTests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.in.Validate(); err.Error() != tt.out {
				t.Errorf("got [%s], want [%s]", err.Error(), tt.out)
			}
		})
	}
}

func TestValidateOK(t *testing.T) {
	c := Convertor{
		Src:  "aaa",
		From: "jpeg",
		To:   "png",
	}
	if err := c.Validate(); err != nil {
		t.Errorf(err.Error())
	}
}
