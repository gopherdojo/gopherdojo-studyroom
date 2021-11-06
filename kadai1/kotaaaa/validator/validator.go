package validator

import (
	"errors"
	"os"
)

func ValidateArgs(targetPath string, targetSrcExt string, targetDstExt string) error {
	if _, err := os.Stat(targetPath); err != nil {
		return errors.New("Error: Doesn't exists the directory that you specified")
	}
	if !validateFileFormat(targetSrcExt) || !validateFileFormat(targetDstExt) {
		return errors.New("Error: Invalid or Unsupported file format")
	}
	return nil
}

func validateFileFormat(target string) bool {
	for _, ext := range []string{".jpg", ".jpeg", ".png", ".gif"} {
		if ext == target {
			return true
		}
	}
	return false
}
