package convert

import (
	"fmt"
	"os"
	"path/filepath"
)

func logError(err error, msg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(1)
	}
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
