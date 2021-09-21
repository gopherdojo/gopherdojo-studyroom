package fileutil

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GetDirName(path string) string {
	return filepath.Dir(path)
}

func GetRelPath(basePath string, targetPath string) (string, error) {
	relPath, err := filepath.Rel(basePath, targetPath)
	if err != nil {
		return "", err
	}
	return relPath, nil
}

func GetFileName(path string) string {
	return filepath.Base(path)
}

func GetFileStem(path string) string {
	pathLength := len(path)
	extLength := len(filepath.Ext(path))
	return filepath.Base(path[:pathLength-extLength])
}

func GetFormattedFileExt(path string) string {
	ext := GetFileExt(path)
	return FormatExt(ext)
}

func GetFileExt(path string) string {
	return filepath.Ext(path)
}

func FormatExt(ext string) string {
	return strings.ToLower(ext)
}

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

func DeleteFile(path string) error {
	return os.Remove(path)
}
