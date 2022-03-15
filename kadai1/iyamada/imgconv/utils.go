package imgconv

import (
	"path/filepath"
	"strings"
)

func isValidFileExtent(ext string) bool {
	switch ext {
	case "jpg", "jpeg", "png", "gif":
		break
	default:
		return false
	}
	return true
}

func hasValidFileExtent(path string, ext string) bool {
	switch ext {
	case "jpg", "jpeg":
		return filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".jpeg"
	default:
		return filepath.Ext(path) == "."+ext
	}
}

func genOutPath(path string, outExt string) (outPath string) {
	return replaceExt(path, outExt)
}

func replaceExt(path, ext string) string {
	return strings.TrimSuffix(path, filepath.Ext(path)) + "." + ext
}
