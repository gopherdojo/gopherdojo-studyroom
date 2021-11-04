package convert

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// User create Type
type fileInfo struct {
	srcFilename string
	dstExt      string
	basePath    string
}

// Create new FileInfo
func NewFileInfo(srcFilename string, dstExt string, basePath string) *fileInfo {
	return &fileInfo{srcFilename, dstExt, basePath}
}

// Remove file
func removeFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	return nil
}

//Get strings except for file's extension.
func getFilePathFromBase(path string) string {
	return path[:len(path)-len(filepath.Ext(path))]
}

// Convert file from src extension to dst extension
func (fi *fileInfo) Convert() error {

	dstFileName := getFilePathFromBase(fi.srcFilename) + fi.dstExt
	// Open target image file object
	srcFile, err := os.Open(fi.basePath + fi.srcFilename)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Read target image file
	img, _, err := image.Decode(srcFile)
	if err != nil {
		return err
	}

	// Create transformed file object
	dstFile, err := os.Create(fi.basePath + dstFileName)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Execute file transform
	switch filepath.Ext(fi.basePath + dstFileName) {
	case ".gif":
		err = gif.Encode(dstFile, img, nil)
	case ".png":
		err = png.Encode(dstFile, img)
	case ".jpg", "jpeg":
		err = jpeg.Encode(dstFile, img, nil)
	default:
		return err
	}
	// Remove src file
	err = removeFile(fi.basePath + fi.srcFilename)

	if err != nil {
		return err
	}
	return nil
}
