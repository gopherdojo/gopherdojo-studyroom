// Package files provides utility functions for files.
package files

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Open a file
func Open(path string) (*os.File, error) {
	return os.Open(filepath.Clean(path))
}

// Get a file's stem (a file name without the extension) by a path
func Stem(path string) string {
	pathLength := len(path)
	extLength := len(filepath.Ext(path))
	return filepath.Base(path[:pathLength-extLength])
}

// Get a formatted file extension by a path
func Ext(path string) string {
	ext := filepath.Ext(path)
	ext = strings.TrimLeft(ext, ".")
	return strings.ToLower(ext)
}

// Get MIME type by a path
func MimeType(path string) (string, error) {
	file, err := Open(path)
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
