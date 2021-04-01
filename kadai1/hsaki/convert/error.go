package convert

import (
	"errors"
	"fmt"
)

var ErrSrcDirPath = errors.New("invalid src directory path")
var ErrDstDirPath = errors.New("invalid dst directory path")
var ErrExt = errors.New("invalid file extension")
var ErrAccessFile = errors.New("cannot access file")
var ErrOpenFile = errors.New("cannot open file")
var ErrCreateImg = errors.New("cannot create image from file")
var ErrOutputPath = errors.New("cannot get output filepath for")
var ErrOutputFile = errors.New("cannot create output file")
var ErrEncodeFile = errors.New("cannot encode img file")

type ConvError struct {
	Err      error
	FilePath string
}

func (e *ConvError) Error() string {
	return fmt.Sprintln(e.Err, e.FilePath)
}

func (e *ConvError) Unwrap() error { return e.Err }
