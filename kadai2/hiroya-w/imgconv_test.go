package imgconv_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/gopherdojo-studyroom/kadai2/hiroya-w/imgconv"
)

func TestEncoder(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		OutputType string
	}{
		{name: "toJPG", OutputType: "jpg"},
		{name: "toPNG", OutputType: "png"},
		{name: "toGIF", OutputType: "gif"},
		{name: "toTIFF", OutputType: "tiff"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &imgconv.Config{
				OutputType: tt.OutputType,
			}
			enc, err := imgconv.NewEncoder(config.OutputType)
			switch tt.OutputType {
			case "jpg":
				if err != nil {
					t.Errorf("NewEncoder() error = %s", err)
				}
				if _, ok := enc.(*imgconv.JPGEncoder); !ok {
					t.Errorf("It is not JPGEncoder. You get %T", enc)
				}
			case "png":
				if err != nil {
					t.Errorf("NewEncoder() error = %s", err)
				}
				if _, ok := enc.(*imgconv.PNGEncoder); !ok {
					t.Errorf("It is not PNGEncoder. You get %T", enc)
				}
			case "gif":
				if err != nil {
					t.Errorf("NewEncoder() error = %s", err)
				}
				if _, ok := enc.(*imgconv.GIFEncoder); !ok {
					t.Errorf("It is not GIFEncoder. You get %T", enc)
				}
			default:
				if err == nil {
					t.Errorf("NewEncoder needs to return an error. But enc = %T", enc)
				}
			}
		})
	}
}

func TestGetFiles(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		inputType string
		directory string
		want      string
	}{
		{name: "jpg", inputType: "jpg", directory: "testdata"},
		{name: "png", inputType: "png", directory: "testdata"},
		{name: "no_such_dir", inputType: "jpg", directory: "hogehoge", want: "no such file or directory"},
		{name: "no_such_dir", inputType: "jpg", directory: "testdata/image1.png", want: "is not a directory"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outStream := new(bytes.Buffer)
			dec, err := imgconv.NewDecoder(tt.inputType)
			if err != nil {
				t.Errorf("NewDecoder() error = %s", err)
			}

			imgConv := &imgconv.ImgConv{
				OutStream: outStream,
				Decoder:   dec,
				TargetDir: tt.directory,
			}
			_, err = imgConv.GetFiles()
			if err != nil {
				if !strings.Contains(err.Error(), tt.want) {
					t.Errorf("expected %q to eq %q", err.Error(), tt.want)
				}
			}
		})
	}
}
