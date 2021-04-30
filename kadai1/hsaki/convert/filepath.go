package convert

import "path/filepath"

// 引数のpathを絶対パスにして返す関数
func absPath(path string) (string, error) {
	if !filepath.IsAbs(path) {
		abspath, err := filepath.Abs(path)
		if err != nil {
			return "", err
		}
		return abspath, nil
	}
	return path, nil
}

// ファイル名から拡張子をとった文字列を返す関数
func removeFileExt(path string) string {
	fileNameLen := len(path)
	extLen := len(filepath.Ext(path))
	return path[:fileNameLen-extLen]
}
