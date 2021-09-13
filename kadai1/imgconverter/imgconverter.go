/*
Package imgconverter provides a method to output an image in a different format.
*/
package imgconverter

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// validInputExtension is a list of image formats that can be specified for the -i option in the argument.
var validInputExtension = []string{
	".jpeg",
	".jpg",
	".png",
	".gif",
}

// validOutputExtension is a list of image formats that can be specified for the -o option in the argument.
// Please note that .gif is not selectable.
var validOutputExtension = []string{
	".jpeg",
	".jpg",
	".png",
}

// IsValidExtension is a function to determine if the extension is selectable.
func IsValidExtension(s string, optionType string) bool {
	var arr []string
	if optionType == "input" {
		arr = validInputExtension
	} else {
		arr = validOutputExtension
	}
	for _, extension := range arr {
		if s == extension {
			return true
		}
	}
	return false
}

// Option stores the values of the -i and -o options entered on the command line.
type Option struct {
	Input, Output *string
}

// Do function transform the image
func Do(path string, option Option) (err error) {
	inputFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		err = inputFile.Close()
	}()

	var img image.Image
	switch *option.Input {
	case ".png":
		img, err = png.Decode(inputFile)
	case ".jpeg", ".jpg":
		img, err = jpeg.Decode(inputFile)
	case ".gif":
		img, err = gif.Decode(inputFile)
	default:
		err = errors.New("invalid input format")
	}
	if err != nil {
		return err
	}

	outputPath := "./icon" + *option.Output
	outputFile, err := os.Create(filepath.Join(filepath.Dir(path), outputPath))
	if err != nil {
		return err
	}
	defer func() {
		err = outputFile.Close()
	}()

	switch *option.Output {
	case ".png":
		err = png.Encode(outputFile, img)
	case ".jpeg", ".jpg":
		err = jpeg.Encode(outputFile, img, nil)
	default:
		err = errors.New("invalid output format")
	}
	if err != nil {
		return err
	}

	return nil
}
