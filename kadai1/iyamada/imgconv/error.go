package imgconv

import (
	"errors"
	"fmt"
)

// errorをまとめた変数
var (
	invalidArg = errors.New("error: invalid argument")
	invalidExt = errors.New("error: invalid extension")
	dirsErr    = errors.New("error: dirs is nil")
)

func wrapError(old, new error) error {
	if old == nil && new == nil {
		return nil
	}
	if old == nil {
		return new
	}
	if new == nil {
		return old
	}
	return fmt.Errorf("%v\n%v", old, new)
}

func wrapErrorWithPath(err error, path string) error {
	return wrapError(err, fmt.Errorf("error: %s is not a valid file", path))
}

func wrapErrorWithTrim(old, new error) error {
	return wrapError(old, fmt.Errorf("error: %s", trimError(new)))
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

func isErrorOccured(err error) bool {
	return err != nil
}
