// imgcnv is a package of functions to convert images
package imgcnv

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

type Extension string

// Convert takes a filename of image and convert it to another imagefile of out
func Convert(filename string, in, out Extension) error {
	// get file extension in lower case (to compare)
	in = Extension(strings.ToLower(string(in)))
	out = Extension(strings.ToLower(string(out)))

	// check input file extension
	switch in {
	case "jpg", "jpeg", "png":
		// nop
	default:
		return errors.New(" invalid input file extension")
	}

	// check output file extension
	switch out {
	case "jpg", "jpeg", "png":
		// nop
	default: // error
		return errors.New(" invalid input file extension")
	}

	// read and decode file
	img, err := readImageFile(filename, in)
	if err != nil {
		return err
	}

	// encode img and write to file
	return writeImageToFile(img, filename+"."+string(out), out)
}

func readImageFile(filename string, fileExt Extension) (image.Image, error) {
	// openfile
	fs, err := os.Open(filename)
	if err != nil {
		return nil, errors.New(" failed to open file")
	}
	defer fs.Close()

	// decode file to Image.image
	if fileExt == "jpg" || fileExt == "jpeg" {
		return jpeg.Decode(fs)
	} else if fileExt == "png" {
		return png.Decode(fs)
	} else {
		return nil, errors.New(" invalid extension")
	}
}

func writeImageToFile(img image.Image, filename string, fileExt Extension) (err error) {
	// openfile
	fs, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() {
		err = fs.Close()
	}()

	// encode and write file
	if fileExt == "jpg" || fileExt == "jpeg" {
		return jpeg.Encode(fs, img, nil)
	} else if fileExt == "png" {
		return png.Encode(fs, img)
	} else {
		return errors.New(" invalid extension")
	}
}
