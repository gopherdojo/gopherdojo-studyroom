package imgconv_test

import (
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
