package picconvert_test

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	picconvert "github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai2/Mizushima/picconvert"
)

func TestPicConverter_Conv(t *testing.T) {
	tmpTestDir := "../tmp_test_dir"
	testCopyDir(t, "../testdata", tmpTestDir)
	testCopyDir(t, "../testdata", tmpTestDir+"/"+filepath.Base(tmpTestDir))

	cases := []struct {
		name string
		pc *picconvert.PicConverter
		expected error
	}{
		{name: "pre: jpeg, after: png", pc: picconvert.NewPicConverter(tmpTestDir, "jpeg", "png"), expected: nil},
		{name: "pre: jpeg, after: gif", pc: picconvert.NewPicConverter(tmpTestDir,"jpeg", "gif"), expected: nil},
		{name: "pre: jpg, after: png", pc: picconvert.NewPicConverter(tmpTestDir,"jpg", "png"), expected: nil},
		{name: "pre: jpg, after: gif", pc: picconvert.NewPicConverter(tmpTestDir, "jpg", "gif"), expected: nil},
		{name: "pre: png, after: jpg", pc: picconvert.NewPicConverter(tmpTestDir,"png", "jpg"), expected: nil},
		{name: "pre: png, after: gif", pc: picconvert.NewPicConverter(tmpTestDir, "png", "gif"), expected: nil},
		{name: "pre: gif, after: jpg", pc: picconvert.NewPicConverter(tmpTestDir, "gif", "jpg"), expected: nil},
		{name: "pre: gif, after: png", pc: picconvert.NewPicConverter(tmpTestDir, "gif", "png"), expected: nil},
		{name: "pre: jpg, after: xls", pc: picconvert.NewPicConverter(tmpTestDir, "jpg", "xls"), expected: fmt.Errorf("xls is not supported")},
		{name: "no file", pc: picconvert.NewPicConverter(tmpTestDir+"/test02.png", "gif", "png"), expected: fmt.Errorf("there's no [gif] file")},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			actual := c.pc.Conv();
			// if actual != nil { fmt.Println(actual) }
			if actual != nil && actual.Error() != c.expected.Error() {
				t.Errorf("want p.Conv() = %v, got %v",c.expected, actual)
			} else if _, err := os.Stat(fmt.Sprintf("%s/test01_converted.%s", tmpTestDir, c.pc.AfterFormat)); err != nil {
				t.Errorf("%s file wasn't made", c.pc.AfterFormat)
			}
		})}

	defer func() {
		err := os.RemoveAll(tmpTestDir)
		if err != nil {
			t.Fatal(err)
		}
	}()

}

func testCopyDir(t *testing.T, srcDir, distDir string) {
	t.Helper()

	src, err := os.Open(srcDir)
	if err != nil {
		t.Fatal(err)
	}
	defer src.Close()

	// make the directory for copied files.
	if err := os.Mkdir(distDir, 0777); err != nil {
		t.Fatal(err)
	}

	// copy each files in srcDir.
	infos, err := src.ReadDir(-1)
	if err != nil {
		t.Fatal(err)
	}

	for _, info := range infos {
		// fmt.Printf("%s :%t\n", info.Name(), info.IsDir())
		distpath := fmt.Sprintf("%s/%s", distDir, info.Name())
		if !info.IsDir() {
			srcpath := fmt.Sprintf("%s/%s", srcDir, info.Name())
			if err = copyFile(srcpath, distpath); err != nil {
				t.Fatal(err)
			} 
		} else {
			if err := os.Mkdir(distpath, 0777); err != nil {
				t.Fatal(err)
			}
		}
	}
}

// copy the file.
func copyFile(srcfile, distfile string) error {
	src, err := os.Open(srcfile)
	if err != nil {
		return err
	}
	defer src.Close()

	dist, err := os.Create(distfile)
	if err != nil {
		return err
	}
	defer dist.Close()

	if _, err := io.Copy(dist, src); err != nil {
		return err
	}
	
	err = dist.Sync()
	if err != nil {
		return err
	}

	return nil
}
