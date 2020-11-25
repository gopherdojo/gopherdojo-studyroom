package imgconv

import (
	"testing"
)

var validatetestdata = []struct {
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
		"not exist dir",
		Convertor{
			Src: "aaa",
		},
		"aaa directory does not exist",
	},
	{
		"invalid -from",
		Convertor{
			Src:  "../testdata",
			From: "jjpeg",
		},
		".jjpeg is not supported",
	},
	{
		"invalid -to",
		Convertor{
			Src:  "../testdata",
			From: "jpeg",
			To:   "ppng",
		},
		".ppng is not supported",
	},
	{
		"-from and -to are same",
		Convertor{
			Src:  "../testdata",
			From: "png",
			To:   "png",
		},
		"-from and -to are same. -from .png, -to .png",
	},
	{
		"both -from and -to are jpg",
		Convertor{
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
	c := Convertor{
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
	in   Convertor
	out  error
}{
	{
		"no target files",
		Convertor{
			Src:     "../testdata/empty",
			From:    ".jpg",
			To:      ".png",
			IsDebug: true,
		},
		nil,
	},
	{
		"jpg -> png",
		Convertor{
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
		Convertor{
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
		Convertor{
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
