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
		t.Fatalf("expected=%d, actual=%d", status, ExitCodeError)
	}

	expected := "flag provided but not defined: -foo"
	if !strings.Contains(errStream.String(), expected) {
		t.Fatalf("expected %s to contain %s", errStream.String(), expected)
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
		b := checkFormat(tt.ext)
		if b != tt.expect {
			t.Fatal("hogehoge")
		}
	}
}
