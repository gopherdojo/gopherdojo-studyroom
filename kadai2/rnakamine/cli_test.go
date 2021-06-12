package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun_parseError(t *testing.T) {
	inStream := bytes.NewBufferString("")
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{inStream: inStream, outStream: outStream, errStream: errStream}
	args := strings.Split("imgconv -foo bar", " ")

	status := cli.Run(args)
	if status != ExitCodeError {
		t.Fatalf("ExitStatus=%d, want %d", status, ExitCodeError)
	}

	expect := "flag provided but not defined: -foo"
	if !strings.Contains(errStream.String(), expect) {
		t.Fatalf("Output=%s, want %s", errStream.String(), expect)
	}
}

func TestCheckFormat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		ext    string
		expect bool
	}{
		{ext: "jpg", expect: true},
		{ext: "jpeg", expect: true},
		{ext: "png", expect: true},
		{ext: "gif", expect: true},
		{ext: "txt", expect: false},
	}

	for _, tt := range tests {
		truth := checkFormat(tt.ext)
		if truth != tt.expect {
			t.Fatalf(`checkFormat("%s") = %t, want %t`, tt.ext, truth, tt.expect)
		}
	}
}
