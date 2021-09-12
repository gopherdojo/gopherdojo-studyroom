/*
Package imgconverter provides a method to output an image in a different format.
*/
package imgconverter

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// ValidInputExtension is a list of image formats that can be specified for the -i option in the argument.
var ValidInputExtension = []string{
	".jpeg",
	".jpg",
	".png",
	".gif",
}

// ValidOutputExtension is a list of image formats that can be specified for the -o option in the argument.
// Please note that .gif is not selectable.
var ValidOutputExtension = []string{
	".jpeg",
	".jpg",
	".png",
}

// Option stores the values of the -i and -o options entered on the command line.
type Option struct {
	Input, Output *string
}

func Do(path string, option Option) error {
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
	}
	if err != nil {
		return err
	}

	return nil
}
