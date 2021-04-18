package download_test

import (
	"runtime"
	"testing"

	"github.com/kimuson13/gopherdojo-studyroom/kimuson13/download"
)

var testurl string = "https://www.naoshima.net/wp-content/uploads/2019/07/393d0895747d5a947ad3acc35eb09688.pdf"

var options download.Options

var fileName string = "paralleldownload"

func TestParse(t *testing.T) {
	cases := []struct {
		name      string
		args      []string
		eParallel int
		eTimeout  int
		eFilename string
	}{
		{name: "noOption", args: []string{testurl}, eParallel: runtime.NumCPU(), eTimeout: 30, eFilename: fileName},
		{name: "parallelOption", args: []string{"-p=6", testurl}, eParallel: 6, eTimeout: 30, eFilename: fileName},
		{name: "timeoutOption", args: []string{"-t=10", testurl}, eParallel: runtime.NumCPU(), eTimeout: 10, eFilename: fileName},
		{name: "filenameOption", args: []string{"-f=test", testurl}, eParallel: runtime.NumCPU(), eTimeout: 30, eFilename: "test"},
		{name: "PandT", args: []string{"-p=6", "-t=20", testurl}, eParallel: 6, eTimeout: 20, eFilename: fileName},
		{name: "PandF", args: []string{"-p=6", "-f=test", testurl}, eParallel: 6, eTimeout: 30, eFilename: "test"},
		{name: "TandF", args: []string{"-t=20", "-f=test", testurl}, eParallel: runtime.NumCPU(), eTimeout: 20, eFilename: "test"},
		{name: "AllOption", args: []string{"-p=6", "-t=20", "-f=test", testurl}, eParallel: 6, eTimeout: 20, eFilename: "test"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			testParse(t, c.args, c.eParallel, c.eTimeout, c.eFilename)
		})
	}
}

func testParse(t *testing.T, args []string, parallel, timeout int, filename string) {
	t.Helper()
	opt, err := options.Parse(args...)
	if err != nil {
		t.Fatal(err)
	}

	if opt.Parallel != parallel {
		t.Errorf("want %v, got %v", parallel, opt.Parallel)
	}

	if opt.Timeout != timeout {
		t.Errorf("want %v, got %v", timeout, opt.Timeout)
	}

	if opt.Filename != filename {
		t.Errorf("want %v, got %v", filename, opt.Filename)
	}
}
