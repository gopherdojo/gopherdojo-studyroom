package validator

import (
	"errors"
	"os"
)

func ValidateArgs(targetPath string, targetSrcExt string, targetDstExt string) error {
	if _, err := os.Stat(targetPath); err != nil {
		return errors.New("Error: Doesn't exists the directory that you specified")
	}
	if !validateFileFormat(targetSrcExt) {
		return errors.New("Error: Conversion source extention is invalid or unsupported.: " + targetSrcExt)
	}
	if !validateFileFormat(targetDstExt) {
		return errors.New("Error: Conversion destination extention is invalid or unsupported.: " + targetDstExt)
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
