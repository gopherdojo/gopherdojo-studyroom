package imgconv

import (
	"log"
	"os"
	"testing"

	"github.com/otiai10/copy"
)

var tmpTestDirectory = "tmp-test-dir"

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
				{FromPath: "tmp-test-dir/A.jpg", ToPath: "tmp-test-dir/A.png"},
				{FromPath: "tmp-test-dir/B.jpg", ToPath: "tmp-test-dir/B.png"},
			},
		},
		{
			from: "png",
			to:   "jpg",
			expect: []ConvertImage{
				{FromPath: "tmp-test-dir/C.png", ToPath: "tmp-test-dir/C.jpg"},
				{FromPath: "tmp-test-dir/D.png", ToPath: "tmp-test-dir/D.jpg"},
			},
		},
		{
			from: "gif",
			to:   "png",
			expect: []ConvertImage{
				{FromPath: "tmp-test-dir/sub/E.gif", ToPath: "tmp-test-dir/sub/E.png"},
				{FromPath: "tmp-test-dir/sub/F.gif", ToPath: "tmp-test-dir/sub/F.png"},
			},
		},
	}

	for _, tt := range tests {
		images, _ := GetConvertImages(tmpTestDirectory, tt.from, tt.to)
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
			convertImage: ConvertImage{FromPath: "tmp-test-dir/A.jpg", ToPath: "tmp-test-dir/A.png"},
			deleteOption: true,
			fromExist:    false,
			toExist:      true,
		},
		{
			convertImage: ConvertImage{FromPath: "tmp-test-dir/C.png", ToPath: "tmp-test-dir/C.jpg"},
			deleteOption: false,
			fromExist:    true,
			toExist:      true,
		},
		{
			convertImage: ConvertImage{FromPath: "tmp-test-dir/sub/E.gif", ToPath: "tmp-test-dir/sub/E.png"},
			deleteOption: true,
			fromExist:    false,
			toExist:      true,
		},
	}

	for _, tt := range tests {
		tt.convertImage.Convert(tt.deleteOption)

		fromExist := fileExist(tt.convertImage.FromPath)
		if fromExist != tt.fromExist {
			t.Errorf("fromExist=%t, want %t", fromExist, tt.fromExist)
		}

		toExist := fileExist(tt.convertImage.ToPath)
		if toExist != tt.toExist {
			t.Errorf("toExist=%t, want %t", toExist, tt.toExist)
		}
	}
}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
