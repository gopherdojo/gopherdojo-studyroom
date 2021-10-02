/*
Package convert is for converting image formats.
*/
package convert

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

/*
struct ConvertImage is for managing image information.
*/
type ConvertImage struct {
	InputFormat     string
	OutputFormat    string
	InputDirectory  string
	OutputDirectory string
}

/*
func (img *ConvertImage) Do() is for converting.
*/
func (img *ConvertImage) Do() {
	_, err := os.Stat(img.OutputDirectory)
	if os.IsNotExist(err) {
		err := os.Mkdir(img.OutputDirectory, 0777)
		logError(err, "cannot create directory")
	}
	var filecount int
	error := filepath.Walk(img.InputDirectory, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+img.InputFormat {
			filecount++
			file, err := os.Open(path)
			logError(err, "Cannot open file")
			defer file.Close()
			convertImage, _, err := image.Decode(file)
			logError(err, "Failed to convert file to image")
			if img.OutputFormat == format.Png {
				dstPath := filepath.Join(img.OutputDirectory, getFileNameWithoutExt(path)+".png")
				out, err := os.Create(dstPath)
				logError(err, "Failed to create destination path")
				defer out.Close()
				png.Encode(out, convertImage)
			} else if img.OutputFormat == format.Jpg || img.OutputFormat == format.Jpeg {
				dstPath := filepath.Join(img.OutputDirectory, getFileNameWithoutExt(path)+".jpg")
				out, err := os.Create(dstPath)
				logError(err, "Failed to create destination path")
				defer out.Close()
				jpeg.Encode(out, convertImage, nil)
			}
		}
		if err != nil {
			return err
		}
		return nil
	})
	if filecount == 0 {
		fmt.Fprintf(os.Stderr, "Images of %v was not found in %v directory\n", img.InputFormat, img.InputDirectory)
		os.Exit(1)
	}
	logError(error, "Error on filepath.Walk")
	fmt.Println("Succuessfully convert image files")
	fmt.Printf("Check %s directory\n", img.OutputDirectory)
}
