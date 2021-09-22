// Package fileutil provides utility functions for files.
package fileutil

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Get a file's stem (a file name without the extension) by a path
func GetFileStem(path string) string {
	pathLength := len(path)
	extLength := len(filepath.Ext(path))
	return filepath.Base(path[:pathLength-extLength])
}

// Get a formatted file extension by a path
func GetFormattedFileExt(path string) string {
	ext := filepath.Ext(path)
	return FormatFileExt(ext)
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
