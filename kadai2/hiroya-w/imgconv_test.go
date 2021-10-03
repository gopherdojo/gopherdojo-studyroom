package imgconv_test

import (
	"bytes"
	"testing"

	"github.com/gopherdojo-studyroom/kadai2/hiroya-w/imgconv"
)

func TestConverter(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		InputType  string
		OutputType string
	}{
		{name: "JPGtoPNG", InputType: "jpg", OutputType: "png"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outStream := new(bytes.Buffer)
			imgConv := &imgconv.ImgConv{
				OutStream: outStream,
			}
			config := &imgconv.Config{
				InputType:  tt.InputType,
				OutputType: tt.OutputType,
			}
			converter := imgconv.NewConverter(config)
			imgConv.Run(converter, "testdata")
		})
	}
}
