package convert

import "path/filepath"

func absPath(path string) string {
	if !filepath.IsAbs(path) {
		abspath, _ := filepath.Abs(path)
		return abspath
	}
	return path
}

func removeFileExt(path string) string {
	fileNameLen := len(path)
	extLen := len(filepath.Ext(path))
	return path[:fileNameLen-extLen]
}
