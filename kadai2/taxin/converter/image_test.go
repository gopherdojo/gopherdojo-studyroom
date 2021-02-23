package converter

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"testing"
)

var (
	x       = 0
	y       = 0
	width   = 400
	height  = 300
	quality = 100
)

func TestFilePathConvert(t *testing.T) {
	filePaths := []struct {
		caseName           string
		filePath           string
		imgFormat          string
		convertedImgFormat string
		convertedFilePath  string
	}{
		{"case1", "./testdata/xx.png", "png", "jpg", "testdata/xx.jpg"},
		{"case2", "./testdata/xx.png.png", "png", "jpg", "testdata/xx.png.jpg"},
		{"case3", "./png/xx.png", "png", "jpg", "png/xx.jpg"},
		{"case4", "testdata/xx.png", "png", "jpg", "testdata/xx.jpg"},
		{"case5", "testdata/xx.jpg", "jpg", "png", "testdata/xx.png"},
		{"case6", "testdata/xx.jpg", "jpg", "gif", "testdata/xx.gif"},
		{"case7", "testdata/xx.jpeg", "jpeg", "png", "testdata/xx.png"},
	}

	for _, tt := range filePaths {
		t.Run(tt.caseName, func(t *testing.T) {
			got := filePathConvert(tt.filePath, tt.imgFormat, tt.convertedImgFormat)
			want := tt.convertedFilePath

			if got != want {
				t.Errorf("conversion of file path failed / got: %q / want: %q", got, want)
			}
		})
	}
}

func TestConvertImgFile(t *testing.T) {
	testCases := []struct {
		caseName   string
		filePath   string
		imgData    ImgDirData
		fileFormat string
	}{
		{"case1", "../testdata/test1.png", ImgDirData{"../testdata", "png", "gif"}, "image/gif"},
		{"case2", "../testdata/test2.gif", ImgDirData{"../testdata", "gif", "jpeg"}, "image/jpeg"},
		{"case3", "../testdata/test3.jpeg", ImgDirData{"../testdata", "jpeg", "png"}, "image/png"},
		{"case4", "../testdata/test4.gif", ImgDirData{"../testdata", "gif", "jpg"}, "image/jpeg"},
	}

	for _, tt := range testCases {
		t.Run(tt.caseName, func(t *testing.T) {

			// create image files for tests
			createImgFileForTesting(t, tt.filePath, tt.imgData)

			// walk dir and convert image
			WalkAndConvertImgFiles(tt.imgData)

			// check the existance of files
			newFilePath := filePathConvert(tt.filePath, tt.imgData.ImgFormat, tt.imgData.ConvertedImgFormat)
			if _, err := os.Stat(newFilePath); err != nil {
				t.Error("does not exists image file that is converted")
			}

			// check whether the file is valid
			// https://golang.org/pkg/net/http/#DetectContentType
			buff := make([]byte, 512)
			srcImgFile, err := os.Open(newFilePath)
			if _, err = srcImgFile.Read(buff); err != nil {
				t.Error("failed to read file header")
			}

			if tt.fileFormat != http.DetectContentType(buff) {
				t.Error("the file is not valid file type")
			}

			// delete test files
			fileErr := os.Remove(tt.filePath)
			newFileErr := os.Remove(newFilePath)
			if fileErr != nil {
				log.Fatal(fileErr)
			}
			if newFileErr != nil {
				log.Fatal(newFileErr)
			}
		})
	}
}

func TestConvertOtherKindsOfFiles(t *testing.T) {
	testCases := []struct {
		caseName      string
		filePath      string
		imgData       ImgDirData
		errExistsFlag bool
	}{
		{"case1", "../testdata/test1.txt", ImgDirData{"../testdata", "txt", "gif"}, true},
		{"case2", "../testdata/test2.png", ImgDirData{"../testdata", "png", "txt"}, true},
	}

	for _, tt := range testCases {
		t.Run(tt.caseName, func(t *testing.T) {

			// create image files for tests
			os.Create(tt.filePath)

			// convert image
			err := WalkAndConvertImgFiles(tt.imgData)

			// delete test files
			fileErr := os.Remove(tt.filePath)
			if fileErr != nil {
				log.Fatal(fileErr)
			}

			if errExists(t, err) != tt.errExistsFlag {
				t.Errorf("error: %#v", err)
			}
		})
	}
}

func createImgFileForTesting(t *testing.T, filepath string, imgData ImgDirData) error {
	t.Helper()

	// create image for testing
	file, err := os.Create(filepath)
	if err != nil {
		t.Error("failed to create file")
	}
	defer file.Close()

	img := image.NewRGBA(image.Rect(x, y, width, height))
	for i := img.Rect.Min.Y; i < img.Rect.Max.Y; i++ {
		for j := img.Rect.Min.X; j < img.Rect.Max.X; j++ {
			img.Set(j, i, color.RGBA{255, 255, 0, 0})
		}
	}

	switch imgData.ImgFormat {
	case "png":
		if err := png.Encode(file, img); err != nil {
			return err
		}
	case "jpeg", "jpg":
		if err := jpeg.Encode(file, img, &jpeg.Options{}); err != nil {
			return err
		}
	case "gif":
		if err := gif.Encode(file, img, nil); err != nil {
			return err
		}
	}

	if err := file.Close(); err != nil {
		return err
	}
	fmt.Println("created files for testing")
	return nil
}

func errExists(t *testing.T, err error) bool {
	t.Helper()
	if err != nil {
		return true
	}
	return false
}
