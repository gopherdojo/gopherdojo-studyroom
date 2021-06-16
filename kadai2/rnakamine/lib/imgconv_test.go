package imgconv

import (
	"log"
	"os"
	"testing"

	"github.com/otiai10/copy"
)

var tmpTestDirectory = "tmp-test-dir/"

func TestMain(m *testing.M) {
	// setup
	if err := copy.Copy("../testdata", tmpTestDirectory); err != nil {
		log.Fatal(err)
	}

	exitCode := m.Run()

	// teardown
	if err := os.RemoveAll(tmpTestDirectory); err != nil {
		log.Fatal(err)
	}

	os.Exit(exitCode)
}

func TestGetConvertImages(t *testing.T) {
	t.Parallel()

	tests := []struct {
		from, to string
		expect   []ConvertImage
	}{
		{
			from: "jpg",
			to:   "png",
			expect: []ConvertImage{
				{FromPath: tmpTestDirectory + "A.jpg", ToPath: tmpTestDirectory + "A.png"},
				{FromPath: tmpTestDirectory + "B.jpg", ToPath: tmpTestDirectory + "B.png"},
			},
		},
		{
			from: "png",
			to:   "jpg",
			expect: []ConvertImage{
				{FromPath: tmpTestDirectory + "C.png", ToPath: tmpTestDirectory + "C.jpg"},
				{FromPath: tmpTestDirectory + "D.png", ToPath: tmpTestDirectory + "D.jpg"},
			},
		},
		{
			from: "gif",
			to:   "png",
			expect: []ConvertImage{
				{FromPath: tmpTestDirectory + "sub/E.gif", ToPath: tmpTestDirectory + "sub/E.png"},
				{FromPath: tmpTestDirectory + "sub/F.gif", ToPath: tmpTestDirectory + "sub/F.png"},
			},
		},
	}

	for _, tt := range tests {
		images, err := GetConvertImages(tmpTestDirectory, tt.from, tt.to)
		if err != nil {
			t.Fatal(err)
		}

		for index, image := range images {
			if image.FromPath != tt.expect[index].FromPath {
				t.Errorf("FromPath=%s, want %s", image.FromPath, tt.expect[index].FromPath)
			}

			if image.ToPath != tt.expect[index].ToPath {
				t.Errorf("ToPath=%s, want %s", image.ToPath, tt.expect[index].ToPath)
			}
		}
	}
}

func TestConvert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		convertImage                     ConvertImage
		deleteOption, fromExist, toExist bool
	}{
		{
			convertImage: ConvertImage{FromPath: tmpTestDirectory + "A.jpg", ToPath: tmpTestDirectory + "A.png"},
			deleteOption: true,
			fromExist:    false,
			toExist:      true,
		},
		{
			convertImage: ConvertImage{FromPath: tmpTestDirectory + "C.png", ToPath: tmpTestDirectory + "C.jpg"},
			deleteOption: false,
			fromExist:    true,
			toExist:      true,
		},
		{
			convertImage: ConvertImage{FromPath: tmpTestDirectory + "sub/E.gif", ToPath: tmpTestDirectory + "sub/E.png"},
			deleteOption: true,
			fromExist:    false,
			toExist:      true,
		},
	}

	for _, tt := range tests {
		err := tt.convertImage.Convert(tt.deleteOption)
		if err != nil {
			t.Fatal(err)
		}

		fromExist := testFileExist(t, tt.convertImage.FromPath)
		if fromExist != tt.fromExist {
			t.Errorf("fromExist=%t, want %t", fromExist, tt.fromExist)
		}

		toExist := testFileExist(t, tt.convertImage.ToPath)
		if toExist != tt.toExist {
			t.Errorf("toExist=%t, want %t", toExist, tt.toExist)
		}
	}
}

func testFileExist(t *testing.T, filename string) bool {
	t.Helper()
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
