package eimg

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

// TestEimg tests functions in eimg package as unittest.
func TestEimg(t *testing.T) {
	TSetParameters(t)
	TEncodeFile(t)
	TConvertExtension(t)
}

// TSetParameters tests SetPerameters().
func TSetParameters(t *testing.T) {

	cases := []struct {
		name     string
		rootDir  string
		fromExt  string
		toExt    string
		expected []string
	}{
		{name: "set RootDir only", rootDir: "test/documents", fromExt: "", toExt: "", expected: []string{"test/documents", "jpeg", "png"}},
		{name: "set RootDir and FromExt", rootDir: "test/img", fromExt: "gif", toExt: "", expected: []string{"test/img", "gif", "png"}},
		{name: "set RootDir and ToExt", rootDir: "test/img", fromExt: "", toExt: "gif", expected: []string{"test/img", "jpeg", "gif"}},
		{name: "set all arguments", rootDir: "test/img", fromExt: "gif", toExt: "jpeg", expected: []string{"test/img", "gif", "jpeg"}},
	}

	unzip := exec.Command("unzip", "test.zip")
	if err := unzip.Run(); err != nil {
		t.Errorf("failed to unzip")
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

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)

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

			eimg := New()
			if err := eimg.SetParameters(); err != nil {
				t.Errorf("failed to set parameter: %#v", err)
			}

			if eimg.RootDir != c.expected[0] {
				t.Errorf("RootDir=> Actual: %s, Expected: %s", eimg.RootDir, c.expected[0])

			}
			if eimg.FromExt != c.expected[1] {
				t.Errorf("FromExt=> Actual: %s, Expected: %s", eimg.FromExt, c.expected[1])

			}

			if eimg.ToExt != c.expected[2] {
				t.Errorf("ToExt=> Actual: %s, Expected: %s", eimg.ToExt, c.expected[2])

			}

			os.Args = []string{}
		})

	}
}

// TEncodeFile tests EncodeFile()
func TEncodeFile(t *testing.T) {
	cases := []struct {
		name     string
		filePath string
		fromExt  string
		toExt    string
		expected string
	}{
		{name: "check png", filePath: "test/img/green.jpeg", fromExt: "jpeg", toExt: "png", expected: "test/img/green.png"},
		{name: "check jpg", filePath: "test/img/blue.gif", fromExt: "gif", toExt: "jpeg", expected: "test/img/blue.jpeg"},
		{name: "check gif", filePath: "test/img/red.png", fromExt: "png", toExt: "gif", expected: "test/img/red.gif"},
	}

	eimg := New()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)
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

			eimg.FromExt = c.fromExt
			eimg.ToExt = c.toExt
			if err := eimg.EncodeFile(c.filePath); err != nil {
				t.Errorf("%s: %s", c.filePath, err)
			}

			if _, err := os.Stat(c.expected); err != nil {
				t.Errorf("%s: %s", c.expected, err)
			}
		})
	}
}

// TConvertExtension tests ConvertExtension()
func TConvertExtension(t *testing.T) {
	cases := []struct {
		name     string
		rootDir  string
		fromExt  string
		toExt    string
		expected []string
	}{
		{name: "check png", rootDir: "test", fromExt: "jpeg", toExt: "png", expected: []string{"test/img/green.png"}},
		{name: "check jpg", rootDir: "test", fromExt: "gif", toExt: "jpeg", expected: []string{"test/img/blue.jpeg"}},
		{name: "check gif", rootDir: "test", fromExt: "png", toExt: "gif", expected: []string{"test/img/red.gif", "test/white.gif"}},
	}

	eimg := New()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)
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
			eimg.RootDir = c.rootDir
			eimg.FromExt = c.fromExt
			eimg.ToExt = c.toExt

			if err := eimg.ConvertExtension(); err != nil {
				t.Errorf("failed ConvertExtension: %s", err)
			}

			for _, filePath := range c.expected {
				if _, err := os.Stat(filePath); err != nil {
					t.Errorf("%s: %s", filePath, err)
				}
			}
		})
	}
}
