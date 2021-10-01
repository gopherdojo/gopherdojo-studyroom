package imgconv_test

import (
	"os"
	"testing"

	"github.com/gopherdojo-studyroom/kadai2/hiroya-w/imgconv"
)

func TestCLI(t *testing.T) {
	cli := &imgconv.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	exitStatus := cli.Run()
	if exitStatus != 0 {
		t.Errorf("Exit status is %d, want %d", exitStatus, 0)
	}
}
