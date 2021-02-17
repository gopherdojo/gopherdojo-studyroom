// Package converter implements conversion of image files
package converter

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ImgDirData defines parameters for converting image files.
type ImgDirData struct {
	DirPath            string
	ImgFormat          string
	ConvertedImgFormat string
}

// WalkAndConvertImgFiles is the function that walks diretory that is passed as args and convert image file format
func WalkAndConvertImgFiles(imgData ImgDirData) {
	err := filepath.Walk(
		imgData.DirPath,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == "."+imgData.ImgFormat {
				convertImgFile(path, imgData)
			}
			return nil
		})
	if err != nil {
		log.Fatal(err)
		return
	}
}

// convertImgFile is the function that checks the typd of image files and convert it to another file type
func convertImgFile(filePath string, imgData ImgDirData) {
	srcImg, decodeErr := decodeImgFile(filePath, imgData)
	if decodeErr != nil {
		fmt.Printf("failed to decode image file: %#v", decodeErr)
	}
	encodeErr := encodeImgFile(filePath, imgData, srcImg)
	if encodeErr != nil {
		fmt.Printf("failed to encode image file: %#v", decodeErr)
	}
}

// decodeImgFile checks files type and decodes the file as an image.Image.
func decodeImgFile(filepath string, imgData ImgDirData) (image.Image, error) {
	fmt.Println("Start to covert image file...")
	srcImgFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer srcImgFile.Close()

	switch imgData.ImgFormat {
	case "png":
		srcImg, err := png.Decode(srcImgFile)
		if err != nil {
			return nil, err
		}
		return srcImg, nil
	case "jpeg", "jpg":
		srcImg, err := jpeg.Decode(srcImgFile)
		if err != nil {
			return nil, err
		}
		return srcImg, nil
	case "gif":
		srcImg, err := gif.Decode(srcImgFile)
		if err != nil {
			return nil, err
		}
		return srcImg, nil
	}
	return nil, errors.New("imgData.ImgFormat is not valid file format")
}

// encodeImgFile checks files type and encodes the file from the data(image.Image)
func encodeImgFile(filePath string, imgData ImgDirData, srcImg image.Image) error {
	file, err := os.Create(filePathConvert(filePath, imgData.ImgFormat, imgData.ConvertedImgFormat))
	if err != nil {
		return err
	}
	defer file.Close()

	switch imgData.ConvertedImgFormat {
	case "png":
		if err := png.Encode(file, srcImg); err != nil {
			return err
		}
	case "jpeg", "jpg":
		if err := jpeg.Encode(file, srcImg, &jpeg.Options{}); err != nil {
			return err
		}
	case "gif":
		if err := gif.Encode(file, srcImg, nil); err != nil {
			return err
		}
	}

	if err := file.Close(); err != nil {
		return err
	}
	fmt.Println("Finished to covert image file")
	return nil
}

func filePathConvert(filePath, imgFormat, convertedImgFormat string) string {
	fileName := filepath.Base(filePath)
	trimmedFileName := strings.TrimRight(fileName, imgFormat)
	return filepath.Dir(filePath) + "/" + trimmedFileName + convertedImgFormat
}
