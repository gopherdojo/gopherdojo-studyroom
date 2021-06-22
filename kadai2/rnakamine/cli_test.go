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
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeError)
	}

	expected := "flag provided but not defined: -foo"
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("Output=%s, want %s", errStream.String(), expected)
	}
}

func TestIsValidFormat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		ext      string
		expected bool
	}{
		{ext: "jpg", expected: true},
		{ext: "jpeg", expected: true},
		{ext: "png", expected: true},
		{ext: "gif", expected: true},
		{ext: "txt", expected: false},
	}

	for _, tt := range tests {
		truth := isValidFormat(tt.ext)
		if truth != tt.expected {
			t.Errorf(`checkFormat("%s") = %t, want %t`, tt.ext, truth, tt.expected)
		}
	}
}
