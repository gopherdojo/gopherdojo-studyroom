package convert

import (
	"errors"
	"fmt"
)

type ErrCode string

// ConvErrorのCodeフィールドにいれるエラーたち
var (
	InValidSrcDirPath ErrCode = "invalid src directory path"
	InValidDstDirPath ErrCode = "invalid dst directory path"
	InValidExt        ErrCode = "invalid file extension"
	FileAccessFail    ErrCode = "cannot access file"
	FileOpenFail      ErrCode = "cannot open file"
	ImgCreateFail     ErrCode = "cannot create image from file"
	InValidOutputPath ErrCode = "cannot get output filepath for"
	FileOutputFail    ErrCode = "cannot create output file"
	FileEncodeFail    ErrCode = "cannot encode img file"
)

// ConvErrorのErrフィールドにいれるエラーたち
var (
	ErrExt = errors.New("invalid file extension")
)

// convertパッケージの公開関数・公開メソッドが返すエラーは全てこの形
type ConvError struct {
	Err      error
	Code     ErrCode
	FilePath string
}

func (e *ConvError) Error() string {
	return fmt.Sprintln(e.Code, e.FilePath)
}

func (e *ConvError) Unwrap() error { return e.Err }
