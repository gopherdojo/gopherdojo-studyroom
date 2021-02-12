// Package converter implements conversion of image files
package converter

import (
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
	err := filepath.Walk(imgData.DirPath,
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
	srcImg := decodeImgFile(filePath, imgData)
	encodeImgFile(filePath, imgData, srcImg)
}

// decodeImgFile checks files type and decodes the file as an image.Image.
func decodeImgFile(filepath string, imgData ImgDirData) image.Image {
	fmt.Println("Start to covert image file...")
	srcImgFile, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}
	defer srcImgFile.Close()

	switch imgData.ImgFormat {
	case "png":
		srcImg, err := png.Decode(srcImgFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil
		}
		return srcImg
	case "jpeg", "jpg":
		srcImg, err := jpeg.Decode(srcImgFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil
		}
		return srcImg
	case "gif":
		srcImg, err := gif.Decode(srcImgFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil
		}
		return srcImg
	}
	return nil
}

// encodeImgFile checks files type and encodes the file from the data(image.Image)
func encodeImgFile(filePath string, imgData ImgDirData, srcImg image.Image) {
	file, err := os.Create(strings.Replace(filePath, imgData.ImgFormat, imgData.ConvertedImgFormat, 1))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	switch imgData.ConvertedImgFormat {
	case "png":
		if err := png.Encode(file, srcImg); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	case "jpeg", "jpg":
		if err := jpeg.Encode(file, srcImg, &jpeg.Options{}); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	case "gif":
		if err := gif.Encode(file, srcImg, nil); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Finished to covert image file")
}
