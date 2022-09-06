package convertor

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/fs"
	"os"
	"path/filepath"
)

// Data required for image conversion
type ImgConvertor struct {
	DirPath   string
	BeforeExt string
	AfterExt  string
}

// Constructor for a ImgConverter.
func NewConvertor(dirPath, beforeExt, afterExt string) ImgConvertor {
	return ImgConvertor{dirPath, beforeExt, afterExt}
}

// Convert images according to the specified format
func (c *ImgConvertor) ConvertImage() error {
	filePaths, err := getFilePaths(c.DirPath, c.BeforeExt)
	if err != nil {
		fmt.Println("getFilePaths:", err)
		return err
	}
	if filePaths == nil {
		return fmt.Errorf("no %s file", c.BeforeExt)
	}
	for _, filePath := range filePaths {
		f, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer f.Close()
		img, _, err := image.Decode(f)
		if err != nil {
			return err
		}
		out, err := os.Create(getConvertedFilePath(filePath, c.AfterExt))
		if err != nil {
			fmt.Println("create:", err)
			return err
		}
		defer out.Close()
		switch c.AfterExt {
		case "jpg", "jpeg":
			err = jpeg.Encode(out, img, &jpeg.Options{})
		case "png":
			err = png.Encode(out, img)
		case "gif":
			err = gif.Encode(out, img, nil)
		default:
			err = fmt.Errorf("%s is not supported", c.AfterExt)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// Recursively retrieve files with a specific extension in a directory.
func getFilePaths(dirPath, beforeExt string) ([]string, error) {
	var filePaths []string
	err := filepath.Walk(dirPath,
		func(path string, info fs.FileInfo, err error) error {
			if filepath.Ext(path) == "."+beforeExt {
				filePaths = append(filePaths, path)
			}
			return nil
		})
	return filePaths, err
}

// Get file path after conversion
func getConvertedFilePath(path, to string) string {
	ext := filepath.Ext(path)
	return path[:len(path)-len(ext)+1] + to
}
