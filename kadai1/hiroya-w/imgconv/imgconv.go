/*
	Package imgconv provides image converter functions.
	JPG, PNG, and GIF are supported.
*/
package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

// Converter converts the inputType format images in the directory to outputType format.
func Converter(directory, inputType, outputType string) error {
	imgPaths, err := getFiles(directory, inputType)
	if err != nil {
		return err
	}

	for _, path := range imgPaths {
		if err := convert(path, outputType); err != nil {
			return err
		}
	}
	return nil
}

// getFiles returns a list of file paths in a directory with the file extension specified by inputType.
func getFiles(directory, inputType string) ([]string, error) {
	var imgPaths []string

	if f, err := os.Stat(directory); err != nil {
		return nil, err
	} else if !f.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", directory)
	}

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == "."+inputType {
			imgPaths = append(imgPaths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return imgPaths, nil
}

// convert converts the image at filePath to the outputType format.
func convert(filePath, outputType string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("Error closing file: %s\n", err)
		}
	}()

	img, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	outputPath := renameExt(filePath, outputType)
	output, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := output.Close(); err != nil {
			log.Printf("Error closing file: %s\n", err)
		}
	}()

	switch outputType {
	case "jpg", "jpeg":
		return convertJPG(img, output)
	case "png":
		return convertPNG(img, output)
	case "gif":
		return convertGIF(img, output)
	default:
		return fmt.Errorf("%s is not a supported output type", outputType)
	}
}

// renameExt renames the file extension of the file at filePath to newExt.
func renameExt(filePath, newExt string) string {
	return filePath[:len(filePath)-len(filepath.Ext(filePath))] + "." + newExt
}

// convertJPG converts the image to the JPEG format.
func convertJPG(img image.Image, output *os.File) error {
	if err := jpeg.Encode(output, img, nil); err != nil {
		return err
	}
	return nil
}

// convertJPG converts the image to the PNG format.
func convertPNG(img image.Image, output *os.File) error {
	if err := png.Encode(output, img); err != nil {
		return err
	}
	return nil
}

// convertJPG converts the image to the GIF format.
func convertGIF(img image.Image, output *os.File) error {
	if err := gif.Encode(output, img, nil); err != nil {
		return err
	}
	return nil
}
