package converter

import (
	"testing"
)

var validatetestdata = []struct {
	name string
	in   Converter
	out  string
}{
	{
		"not exist dir",
		Converter{
			Src: "aaa",
		},
		"aaa directory does not exist",
	},
	{
		"invalid -from",
		Converter{
			Src:  "../testdata",
			From: "jjpeg",
		},
		".jjpeg is not supported",
	},
	{
		"invalid -to",
		Converter{
			Src:  "../testdata",
			From: "jpeg",
			To:   "ppng",
		},
		".ppng is not supported",
	},
	{
		"-from and -to are same",
		Converter{
			Src:  "../testdata",
			From: "png",
			To:   "png",
		},
		"-from and -to are same. -from .png, -to .png",
	},
	{
		"both -from and -to are jpg",
		Converter{
			Src:  "../testdata",
			From: "jpeg",
			To:   "jpg",
		},
		"-from and -to are same. -from .jpeg, -to .jpg",
	},
}

func TestValidateNG(t *testing.T) {
	for _, tt := range validatetestdata {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.in.Validate(); err.Error() != tt.out {
				t.Errorf("got [%s], want [%s]", err.Error(), tt.out)
			}
		})
	}
}

func TestValidateOK(t *testing.T) {
	c := Converter{
		Src:  "../testdata",
		From: "jpeg",
		To:   "png",
	}
	if err := c.Validate(); err != nil {
		t.Errorf(err.Error())
	}
}

var doConvertorTest = []struct {
	name string
	in   Converter
	out  error
}{
	{
		"no target files",
		Converter{
			Src:     "../testdata/empty",
			From:    ".jpg",
			To:      ".png",
			IsDebug: true,
		},
		nil,
	},
	{
		"jpg -> png",
		Converter{
			Src:     "../testdata/valid_data",
			Dst:     "../output_JpgToPng",
			From:    ".jpg",
			To:      ".png",
			IsDebug: true,
		},
		nil,
	},
	{
		"png -> gif",
		Converter{
			Src:     "../testdata/valid_data",
			Dst:     "../output_PngToGif",
			From:    ".png",
			To:      ".gif",
			IsDebug: true,
		},
		nil,
	},
	{
		"gif -> jpg",
		Converter{
			Src:     "../testdata/valid_data",
			Dst:     "../output_GIFToJPEG",
			From:    ".gif",
			To:      ".jpeg",
			IsDebug: true,
		},
		nil,
	},
}

func TestDoConvert(t *testing.T) {
	for _, tt := range doConvertorTest {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.in.DoConvert(); err != nil {
				t.Errorf("Unexpected error. %s", err)
			}
		})
	}
}
