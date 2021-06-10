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
