// imageCoverter is a package of functions to convert images
package imageConverter

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

type Extension string

// Convert takes a filename of image and convert it to another imagefile of outputFileExt
func Convert(filename string, inputFileExt Extension, outputFileExt Extension) error {
	// get file extension in lower case (to compare)
	inputFileExt = Extension(strings.ToLower(string(inputFileExt)))
	outputFileExt = Extension(strings.ToLower(string(outputFileExt)))

	// check extension
	if inputFileExt != "jpg" && inputFileExt != "jpeg" && inputFileExt != "png" {
		return errors.New("invalid input file extension")
	} else if outputFileExt != "jpg" && outputFileExt != "jpeg" && outputFileExt != "png" {
		return errors.New("invalid output file extension")
	}

	// read and decode file
	img, err := readImageFile(filename, inputFileExt)
	if err != nil {
		return err
	}

	// encode img and write to file
	return writeImageToFile(img, filename+"."+string(outputFileExt), outputFileExt)
}

func readImageFile(filename string, fileExt Extension) (image.Image, error) {
	// openfile
	fs, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	defer fs.Close()

	// decode file to Image.image
	if fileExt == "jpg" || fileExt == "jpeg" {
		return jpeg.Decode(fs)
	} else if fileExt == "png" {
		return png.Decode(fs)
	} else {
		return nil, errors.New("invalid extension")
	}
}

func writeImageToFile(img image.Image, filename string, fileExt Extension) error {
	// openfile
	fs, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("failed to open file")
	}
	defer fs.Close()

	// encode and write file
	if fileExt == "jpg" || fileExt == "jpeg" {
		return jpeg.Encode(fs, img, nil)
	} else if fileExt == "png" {
		return png.Encode(fs, img)
	} else {
		return errors.New("invalid extension")
	}
}
