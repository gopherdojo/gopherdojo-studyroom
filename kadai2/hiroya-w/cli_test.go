package imgconv_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/gopherdojo-studyroom/kadai2/hiroya-w/imgconv"
)

func TestCLI(t *testing.T) {
	t.Parallel()
	tests := []struct {
		options    string
		exitStatus int
		want       string
	}{
		{options: "", exitStatus: 1, want: "directory is required"},
		{options: "-h", exitStatus: 1, want: "Usage"},
		{options: "-input-type=bmp testdata", exitStatus: 1, want: "invalid input type:"},
		{options: "-output-type=tiff testdata", exitStatus: 1, want: "invalid output type:"},
		{options: "-input-type=jpg -output-type=jpg testdata", exitStatus: 1, want: "input type and output type must be different"},
		{options: "testdata", exitStatus: 0, want: ""},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("Options:'%s'", test.options), func(t *testing.T) {
			outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
			cli := &imgconv.CLI{OutStream: outStream, ErrStream: errStream}

			os.Args = append([]string{os.Args[0]}, strings.Split(test.options, " ")...)
			exitStatus := cli.Run()

			if exitStatus != test.exitStatus {
				t.Errorf("exit status = %d, want %d", exitStatus, test.exitStatus)
			}

			if !strings.Contains(errStream.String(), test.want) {
				t.Errorf("expected %q to eq %q", errStream.String(), test.want)
			}
		})
	}
}
