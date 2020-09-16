package imgconv

import (
	"errors"
	"os"
)

func validateDir(d *string) bool {
	if f, err := os.Stat(*d); os.IsNotExist(err) || !f.IsDir() {
		return false
	}
	return true
}

func validateExtension(te *string) bool {
	for _, e := range extensions {
		if string(e) == *te {
			return true
		}
	}
	return false
}

func validation(id *string, od *string, ie *string, oe *string) error {
	if !validateDir(id) {
		return errors.New("invalid input dir")
	}
	if !validateDir(od) {
		return errors.New("invalid output dir")
	}
	if !validateExtension(ie) {
		return errors.New("invalid extension of input file")
	}
	if !validateExtension(oe) {
		return errors.New("invalid extension of output file")
	}
	return nil
}
