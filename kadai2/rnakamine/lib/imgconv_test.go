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
		expected []ConvertImage
	}{
		{
			from: "jpg",
			to:   "png",
			expected: []ConvertImage{
				{FromPath: tmpTestDirectory + "A.jpg", ToPath: tmpTestDirectory + "A.png"},
				{FromPath: tmpTestDirectory + "B.jpg", ToPath: tmpTestDirectory + "B.png"},
			},
		},
		{
			from: "png",
			to:   "jpg",
			expected: []ConvertImage{
				{FromPath: tmpTestDirectory + "C.png", ToPath: tmpTestDirectory + "C.jpg"},
				{FromPath: tmpTestDirectory + "D.png", ToPath: tmpTestDirectory + "D.jpg"},
			},
		},
		{
			from: "gif",
			to:   "png",
			expected: []ConvertImage{
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
			if image.FromPath != tt.expected[index].FromPath {
				t.Errorf("FromPath=%s, want %s", image.FromPath, tt.expected[index].FromPath)
			}

			if image.ToPath != tt.expected[index].ToPath {
				t.Errorf("ToPath=%s, want %s", image.ToPath, tt.expected[index].ToPath)
			}
		}
	}
}

func TestConvert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		convertImage                                     ConvertImage
		deleteOption, expectedFromExist, expectedToExist bool
	}{
		{
			convertImage:      ConvertImage{FromPath: tmpTestDirectory + "A.jpg", ToPath: tmpTestDirectory + "A.png"},
			deleteOption:      true,
			expectedFromExist: false,
			expectedToExist:   true,
		},
		{
			convertImage:      ConvertImage{FromPath: tmpTestDirectory + "C.png", ToPath: tmpTestDirectory + "C.jpg"},
			deleteOption:      false,
			expectedFromExist: true,
			expectedToExist:   true,
		},
		{
			convertImage:      ConvertImage{FromPath: tmpTestDirectory + "sub/E.gif", ToPath: tmpTestDirectory + "sub/E.png"},
			deleteOption:      true,
			expectedFromExist: false,
			expectedToExist:   true,
		},
	}

	for _, tt := range tests {
		err := tt.convertImage.Convert(tt.deleteOption)
		if err != nil {
			t.Fatal(err)
		}

		expectedFromExist := testFileExist(t, tt.convertImage.FromPath)
		if expectedFromExist != tt.expectedFromExist {
			t.Errorf("expectedFromExist=%t, want %t", expectedFromExist, tt.expectedFromExist)
		}

		expectedToExist := testFileExist(t, tt.convertImage.ToPath)
		if expectedToExist != tt.expectedToExist {
			t.Errorf("expectedToExist=%t, want %t", expectedToExist, tt.expectedToExist)
		}
	}
}

func testFileExist(t *testing.T, filename string) bool {
	t.Helper()
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
