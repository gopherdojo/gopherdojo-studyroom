package conv

import (
	"os"
	"testing"
)

func Test_getFileContentType(t *testing.T) {
	type in struct {
		filepath string
	}
	tests := []struct {
		in  in
		out string
	}{
		{in{"test/originally_jpg.jpg"}, ContentTypeJpeg},
		{in{"test/originally_png.png"}, ContentTypePng},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("Test", func(t *testing.T) {
			file, err := os.Open(tt.in.filepath)
			if err != nil {
				t.Errorf("failed to load a file: %s", tt.in.filepath)

				return
			}
			defer file.Close()

			got, err := getFileContentType(file)
			if err != nil {
				t.Errorf("got Error while executing getFileContentType: %s", err)
			}

			if got != tt.out {
				t.Errorf("got %s want %s", got, tt.out)
			}
		})
	}
}
