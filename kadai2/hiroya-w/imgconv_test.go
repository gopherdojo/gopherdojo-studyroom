package imgconv_test

import (
	"strings"
	"testing"

	imgconv "github.com/Hiroya-W/gopherdojo-studyroom/kadai2/hiroya-w"
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
		{name: "no_such_dir", inputType: "jpg", directory: "cmd/imgconv/main.go", want: "is not a directory"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dec, err := imgconv.NewDecoder(tt.inputType)
			if err != nil {
				t.Errorf("NewDecoder() error = %s", err)
			}

			imgConv := &imgconv.ImgConv{
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

func TestGetFilesCount(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		inputType string
		directory string
		want      int
	}{
		{name: "go_files", inputType: "go", directory: ".", want: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dec := &imgconv.ImageDecoder{
				&imgconv.Extention{tt.inputType},
			}
			imgConv := &imgconv.ImgConv{
				Decoder:   dec,
				TargetDir: tt.directory,
			}
			files, err := imgConv.GetFiles()
			if err != nil {
				t.Errorf("GetFiles() error = %s", err)
			}
			if len(files) != tt.want {
				t.Errorf("expected %d to eq %d", len(files), tt.want)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	// t.Parallel()
	tests := []struct {
		name       string
		inputType  string
		outputType string
		inputFile  string
		want       string
	}{
		{name: "JPGtoPNG", inputType: "jpg", outputType: "png", inputFile: "testdata/image_jpg.jpg", want: "testdata/image_jpg.png"},
		{name: "JPGtoGIF", inputType: "jpg", outputType: "gif", inputFile: "testdata/image_jpg.jpg", want: "testdata/image_jpg.gif"},
		{name: "PNGtoJPG", inputType: "png", outputType: "jpg", inputFile: "testdata/image_png.png", want: "testdata/image_png.jpg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dec, err := imgconv.NewDecoder(tt.inputType)
			if err != nil {
				t.Errorf("NewDecoder() error = %s", err)
			}
			enc, err := imgconv.NewEncoder(tt.outputType)
			if err != nil {
				t.Errorf("NewEncoder() error = %s", err)
			}

			imgConv := &imgconv.ImgConv{
				Decoder: dec,
				Encoder: enc,
			}
			outputPath, err := imgConv.Convert(dec, enc, tt.inputFile)
			if err != nil {
				t.Errorf("Convert() error = %s", err)
			}

			if outputPath != tt.want {
				t.Errorf("expected %q to eq %q", outputPath, tt.want)
			}
		})
	}
}
