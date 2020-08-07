package eimg

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestRun(t *testing.T) {
	cases := []struct {
		name    string
		rootDir string
		fromExt string
		toExt   string
	}{

		{name: "set RootDir only", rootDir: "test/documents", fromExt: "", toExt: ""},
		{name: "set RootDir and FromExt", rootDir: "test/img", fromExt: "gif", toExt: ""},
		{name: "set RootDir and ToExt", rootDir: "test/img", fromExt: "", toExt: "gif"},
		{name: "set all arguments", rootDir: "test/img", fromExt: "gif", toExt: "jpeg"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)

			// set arguments
			os.Args = append(os.Args, "go")
			if c.fromExt != "" {
				os.Args = append(os.Args, "-f="+c.fromExt)
			}
			if c.toExt != "" {
				os.Args = append(os.Args, "-t="+c.toExt)
			}
			if c.rootDir != "" {
				os.Args = append(os.Args, c.rootDir)
			}

			// extract test zip file
			unzip := exec.Command("unzip", "test.zip")
			if err := unzip.Run(); err != nil {
				t.Errorf("failed to unzip...")
			}
			defer func() {
				if _, err := os.Stat("test"); err == nil {
					rmAll := exec.Command("rm", "-rf", "./test")
					if err := rmAll.Run(); err != nil {
						return
					}
				} else {
					return
				}
			}()

			eimg := New()
			if err := eimg.Run(); err != nil {
				t.Errorf("faield to Run: %s", err)
			}
		})
	}
}
