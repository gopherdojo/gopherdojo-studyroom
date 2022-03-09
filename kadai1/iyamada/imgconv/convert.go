// Package imgconv implements image converter
package imgconv

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type myImage image.Image

func writeImage(file io.Writer, img myImage, ext string) (err error) {
	switch ext {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(file, img, nil)
	case ".png":
		err = png.Encode(file, img)
	case ".gif":
		err = gif.Encode(file, img, nil)
	}
	return err
}

func readImage(file io.Reader, ext string) (img myImage, err error) {
	switch ext {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	case ".gif":
		img, err = gif.Decode(file)
	}
	return img, err
}

func convert(inPath string, outPath string) (err error) {
	inFile, err := os.Open(inPath)
	if err != nil {
		return err
	}
	inImg, err := readImage(inFile, filepath.Ext(inPath))
	if err != nil {
		return err
	}
	outFile, err := os.Create(outPath)
	if err != nil {
		return err
	}
	if err := writeImage(outFile, inImg, filepath.Ext(outPath)); err != nil {
		return err
	}
	defer func() {
		err = inFile.Close()
	}()
	defer func() {
		err = outFile.Close()
	}()
	return nil
}

// Convert converts image files that exist in a directory passed as a command line argument.
// The file to be converted is specified by -i.
// The file to be converted is specified by -o as well.
// The image formats supported are jpeg, png, and gif.
// If no image format is specified, jpeg files will be converted to png files.
// Even if the specified directory has subdirectories, image files under the subdirectories will be converted.
// If no directory is passed as an argument, an error will be returned.
// It also returns an error if the appropriate image format is not specified.
// If multiple directories are passed, it will search the directories in the order they are passed.
// Even if a text file or other file not to be converted is found during the search, it will continue to convert other files.
func Convert(dirs []string, inExt, outExt string) (convErr error) {
	convErr = errors.New("")
	for _, dir := range dirs {
		walkErr := filepath.WalkDir(dir, func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() || isValidFileExtent(path, outExt) {
				return nil
			}
			if !isValidFileExtent(path, inExt) {
				convErr = fmt.Errorf("error: %s is not a valid file\n%v", path, convErr)
				return nil
			}
			if err := convert(path, getOutPath(path, outExt)); err != nil {
				convErr = fmt.Errorf("error: %s\n%v", trimError(err), convErr)
				return nil
			}
			return nil
		})
		if walkErr != nil {
			convErr = fmt.Errorf("error: %s\n%v", trimError(walkErr), convErr)
		}
	}
	return convErr
}
