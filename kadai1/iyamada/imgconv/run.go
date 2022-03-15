// Package imgconv implements image converter
package imgconv

import (
	"io/fs"
	"path/filepath"
)

func isNoEntry(err error) bool {
	return isErrorOccured(err)
}

func isSkip(info fs.DirEntry, path, ext string) bool {
	return info.IsDir() || hasValidFileExtent(path, ext)
}

func validateArg(dirs []string, inExt, outExt string) error {
	if dirs == nil {
		return dirsErr
	}
	if !isValidFileExtent(inExt) || !isValidFileExtent(outExt) {
		return invalidExt
	}
	return nil
}

// Run converts image files passed as a command line argument.
// Default is jpg to png format.
// If the image conversion fails, an error is returned.
func Run(dirs []string, inExt, outExt string) (convErr error) {
	if err := validateArg(dirs, inExt, outExt); err != nil {
		return err
	}
	for _, entry := range dirs {
		err := filepath.WalkDir(entry, func(path string, info fs.DirEntry, err error) error {
			if isNoEntry(err) {
				return err
			}
			if isSkip(info, path, outExt) {
				return nil
			}
			if !hasValidFileExtent(path, inExt) {
				convErr = wrapErrorWithPath(convErr, path)
				return nil
			}
			if err := convert(path, outExt); err != nil {
				convErr = wrapErrorWithTrim(convErr, err)
			}
			return nil
		})
		if err != nil {
			convErr = wrapErrorWithTrim(convErr, err)
		}
	}
	return convErr
}
