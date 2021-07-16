package main

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"testing"
)


var testdataPathMap = map[int][]string{
	0: {"../testdata/003", mustGetSize("../testdata/003")},
	1: {"../testdata/z4d4kWk.jpg", mustGetSize("../testdata/z4d4kWk.jpg")},
}

func Test_main(t *testing.T) {
	cases := map[string]struct {
		key int // key for testdataPathMap
	}{
		"case 1" : {
			key : 0,
		},
		"case 2" : {
			key : 1,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			testdataPath := testdataPathMap[c.key][0]
			expected, err := os.Open(testdataPath)
			if err != nil {
				t.Fatal(err)
			}

			defer func() {
				if err := expected.Close(); err != nil {
					t.Fatal(err)
				}
			}()

			opts := &Options{
				Help : false,
				Procs: uint(runtime.NumCPU()),
				Output: "./",
				Tm: 3,
			}

			

		})
	}
}

// mustGetSize returns the size of the file in "path" as a string for "Content-Length" in http header.
func mustGetSize(path string) string {

	fileinfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	return strconv.Itoa(int(fileinfo.Size()))
}