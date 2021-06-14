package imgconv

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	testTmpPath, err := ioutil.TempDir("./", "test-tmp-path")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("testTmpPath: %s\n", testTmpPath)

	exitCode := m.Run()

	if err := os.RemoveAll(testTmpPath); err != nil {
		log.Fatal(err)
	}

	os.Exit(exitCode)
}

func TestConvert(t *testing.T) {
	t.Log("TestConvert")
}

func TestGetConvertImages(t *testing.T) {
	t.Parallel()

	dir := "../testdata"
	tests := []struct {
		from, to string
		expect   []ConvertImage
	}{
		{
			from: "jpg",
			to:   "png",
			expect: []ConvertImage{
				{FromPath: "../testdata/A.jpg", ToPath: "../testdata/A.png"},
				{FromPath: "../testdata/B.jpg", ToPath: "../testdata/B.png"},
			},
		},
		{
			from: "png",
			to:   "jpg",
			expect: []ConvertImage{
				{FromPath: "../testdata/C.png", ToPath: "../testdata/C.jpg"},
				{FromPath: "../testdata/D.png", ToPath: "../testdata/D.jpg"},
			},
		},
		{
			from: "gif",
			to:   "png",
			expect: []ConvertImage{
				{FromPath: "../testdata/sub/E.gif", ToPath: "../testdata/sub/E.png"},
				{FromPath: "../testdata/sub/F.gif", ToPath: "../testdata/sub/F.png"},
			},
		},
	}

	for _, tt := range tests {
		images, _ := GetConvertImages(dir, tt.from, tt.to)
		for index, image := range images {
			if image.FromPath != tt.expect[index].FromPath {
				t.Fatalf("FromPath=%s, want %s", image.FromPath, tt.expect[index].FromPath)
			}
			if image.ToPath != tt.expect[index].ToPath {
				t.Fatalf("ToPath=%s, want %s", image.ToPath, tt.expect[index].ToPath)
			}
		}
	}
}
