package imgconv

import (
	"path/filepath"
	"strings"
)

func isValidFileExtent(path string, ext string) bool {
	switch ext {
	case ".jpg", ".jpeg":
		return strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg")
	default:
		return strings.HasSuffix(path, string(ext))
	}
}

func trimError(err error) string {
	s := err.Error()
	for i, c := range s {
		if c == ' ' {
			return s[i+1:]
		}
	}
	return s
}

func getOutPath(path string, outExt string) (outPath string) {
	return replaceFileExtent(path, filepath.Ext(path), "."+outExt)
}

func replaceFileExtent(filePath string, oldExt, newExt string) string {
	if strings.HasSuffix(filePath, ".jpg") && oldExt == ".jpeg" {
		return replaceSuffix(filePath, ".jpg", string(newExt))
	} else if strings.HasSuffix(filePath, ".jpeg") && oldExt == ".jpg" {
		return replaceSuffix(filePath, ".jpeg", string(newExt))
	}
	return replaceSuffix(filePath, string(oldExt), string(newExt))
}

func replaceSuffix(s, old, new string) string {
	return strings.TrimSuffix(s, old) + new
}
