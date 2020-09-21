package conv

import (
	"errors"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	ContentTypeJpeg = "image/jpeg"
	ContentTypePng  = "image/png"
	ExtensionJpg    = ".jpg"
	ExtensionPng    = ".png"
)

type fileDest struct {
	from *os.File
	to   *os.File
}

// ConvertImages converts an image file with an extension to another specified by "extFrom" and "extTo" in "destDir" directory.
func ConvertImages(destDir, extFrom, extTo string) error {
	if extFrom == extTo {
		return errors.New("specified extensions must be distinct")
	}

	return filepath.Walk(destDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, extFrom) {
			err = convert(path, extFrom, extTo)
		}

		return nil
	})
}

func convert(filepath, extFrom, extTo string) error {
	from, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.Create(strings.TrimSuffix(filepath, extFrom) + extTo)
	if err != nil {
		return err
	}
	defer to.Close()

	switch extFrom {
	case ExtensionJpg:
		return jpegToPng(fileDest{from, to})
	case ExtensionPng:
		return pngToJpeg(fileDest{from, to})
	default:
		return nil
	}
}

func jpegToPng(dest fileDest) error {
	if !isJpeg(dest.from) {
		return errors.New("content type of the original file is not " + ContentTypeJpeg)
	}

	jpegImg, err := jpeg.Decode(dest.from)

	if err != nil {
		return err
	}

	png.Encode(dest.to, jpegImg)
	return nil
}

func pngToJpeg(dest fileDest) error {
	if !isPng(dest.from) {
		return errors.New("content type of the original file is not " + ContentTypePng)
	}

	pngImg, err := png.Decode(dest.from)

	if err != nil {
		return err
	}

	return jpeg.Encode(dest.to, pngImg, nil)
}

func isJpeg(file *os.File) bool {
	contentType, _ := getFileContentType(file)
	return contentType == ContentTypeJpeg
}

func isPng(file *os.File) bool {
	contentType, _ := getFileContentType(file)
	return contentType == ContentTypePng
}

func getFileContentType(file *os.File) (string, error) {
	// Using the first 512 bytes to detect the content type.
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	// Reset the file pointer
	file.Seek(0, io.SeekStart)

	if err != nil {
		return "", err
	}

	return http.DetectContentType(buffer), nil
}
