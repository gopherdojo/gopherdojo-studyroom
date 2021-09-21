package fileutil

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Get a directory name by a path
func GetDirName(path string) string {
	return filepath.Dir(path)
}

// Get a relative path
func GetRelPath(basePath string, targetPath string) (string, error) {
	relPath, err := filepath.Rel(basePath, targetPath)
	if err != nil {
		return "", err
	}
	return relPath, nil
}

// Get a file name by a path
func GetFileName(path string) string {
	return filepath.Base(path)
}

// Get a file's stem (a file name without the extension) by a path
func GetFileStem(path string) string {
	pathLength := len(path)
	extLength := len(filepath.Ext(path))
	return filepath.Base(path[:pathLength-extLength])
}

// Get a formatted file extension by a path
func GetFormattedFileExt(path string) string {
	ext := GetFileExt(path)
	return FormatFileExt(ext)
}

// Get a file extension by a path
func GetFileExt(path string) string {
	return filepath.Ext(path)
}

// Format a file extension
func FormatFileExt(ext string) string {
	return strings.ToLower(ext)
}

// Get MIME type by a path
func GetMimeType(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	buf := make([]byte, 512)
	if _, err := file.Read(buf); err != nil {
		return "", err
	}
	mimeType := http.DetectContentType(buf)
	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}
	return mimeType, nil
}

// Delete a file by a path
func DeleteFile(path string) error {
	return os.Remove(path)
}
