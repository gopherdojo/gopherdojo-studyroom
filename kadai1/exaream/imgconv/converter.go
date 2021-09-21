package imgconv

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"kadai1/fileutil"
)

const (
	// Extensions
	extJpg        = ".jpg"
	extJpeg       = ".jpeg"
	extPng        = ".png"
	extGif        = ".gif"
	defaultSrcExt = extJpg
	defaultDstExt = extPng
	// Directories
	defaultSrcDir = "./testdata/src"
	defaultDstDir = "./testdata/dst"
	// Flag
	defaultFileDeleteFlag = false
)

type Converter struct {
	srcExt         string
	dstExt         string
	srcDir         string
	dstDir         string
	fileDeleteFlag bool
}

var (
	// Extensions' list
	extList    []string = []string{extJpg, extJpeg, extPng, extGif}
	extListStr string   = strings.Join(extList, " ")
	// Arguments
	SrcExt         *string = flag.String("src-ext", defaultSrcExt, "Source extension (choices "+extListStr+")")
	DstExt         *string = flag.String("dst-ext", defaultDstExt, "Destination extension (choices "+extListStr+")")
	SrcDir         *string = flag.String("src-dir", defaultSrcDir, "Source directory")
	DstDir         *string = flag.String("dst-dir", defaultDstDir, "Destination directory")
	FileDeleteFlag *bool   = flag.Bool("delete", defaultFileDeleteFlag, "File delete flag")
)

func ValidArgs() error {
	if !containsStringInSlice(extList, *SrcExt) {
		return fmt.Errorf("The selected \"src-ext\" does not exist in: %v", extListStr)
	}

	if !containsStringInSlice(extList, *DstExt) {
		return fmt.Errorf("The selected \"dst-ext\" does not exist in: %v", extListStr)
	}

	dir, err := os.Open(*SrcDir)
	if err != nil {
		return fmt.Errorf("Faild to open the directory of the \"src-dir\": %w", err)
	}

	dirInfo, err := dir.Stat()
	if err != nil {
		return fmt.Errorf("Failed to get the directory info of the \"src-dir\": %w", err)
	}

	if !dirInfo.IsDir() {
		return errors.New("The \"src-dir\" must be an existing directory")
	}

	if err := dir.Close(); err != nil {
		return fmt.Errorf("Failed to close the directory of \"src-dir\": %w", err)
	}

	if getType(*FileDeleteFlag) != "bool" {
		return errors.New("The \"delete\" option must be true or false")
	}
	return nil
}

func NewConverter(srcExt string, dstExt string, srcDir string, dstDir string, fileDeleteFlag bool) *Converter {
	return &Converter{
		srcExt:         srcExt,
		dstExt:         dstExt,
		srcDir:         srcDir,
		dstDir:         dstDir,
		fileDeleteFlag: fileDeleteFlag,
	}
}

func (conv *Converter) Run() error {
	return filepath.Walk(conv.srcDir, func(srcFilePath string, info os.FileInfo, err error) error {
		// filepath.Walk() 内部で os.FileInfo 取得時にエラーがある場合、処理をスキップ
		if err != nil {
			return err
		}
		// ディレクトリの場合、処理をスキップ
		if info.IsDir() {
			return nil
		}
		// 取得した拡張子と変換元の拡張子が異なる場合、処理をスキップ
		if ext := fileutil.GetFormattedFileExt(srcFilePath); ext != conv.srcExt {
			return nil
		}
		// 変換後のディレクトリ・パスを取得
		dstDir, err := conv.getDstDir(srcFilePath)
		if err != nil {
			return err
		}
		// 変換後のディレクトリが存在しない場合は作成 (os.MkdirAll() 内でディレクトリの有無を判定)
		if err = os.MkdirAll(dstDir, 0777); err != nil {
			return err
		}
		// 変換後のファイル・パスを取得
		dstFilePath := conv.getDstFilePath(srcFilePath, dstDir)
		// 画像の変換処理
		if err := conv.convert(srcFilePath, dstFilePath); err != nil {
			return err
		}
		// 変換元の画像を削除
		if conv.fileDeleteFlag {
			if err := fileutil.DeleteFile(srcFilePath); err != nil {
				return err
			}
		}
		return nil
	})
}

func (conv *Converter) getDstDir(srcFilePath string) (string, error) {
	// 変換元の対象ディレクトリの相対パスを取得
	srcRelDir, err := conv.getSrcRelDir(srcFilePath)
	if err != nil {
		return "", err
	}
	// 上記の相対パスと変換後の基準ディレクトリを元に
	// 変換後のディレクトリの絶対パスを取得
	path := filepath.Join(conv.dstDir, srcRelDir)
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return absPath, nil
}

func (conv *Converter) getSrcRelDir(srcFilePath string) (string, error) {
	dir := fileutil.GetDirName(srcFilePath)
	return fileutil.GetRelPath(conv.srcDir, dir)
}

func (conv *Converter) getDstFilePath(srcFilePath string, dstDir string) string {
	dstFileName := fileutil.GetFileStem(srcFilePath) + conv.dstExt
	return filepath.Join(dstDir, dstFileName)
}

func (conv *Converter) convert(srcFilePath string, dstFilePath string) error {
	srcImage, err := getSrcImage(srcFilePath)
	if err != nil {
		return err
	}
	dstImage, err := getDstImage(dstFilePath)
	if err != nil {
		return err
	}
	if err := conv.encode(dstImage, srcImage); err != nil {
		return err
	}
	return nil
}

func (conv *Converter) encode(dstImage io.Writer, srcImage image.Image) error {
	switch conv.dstExt {
	case extPng:
		return png.Encode(dstImage, srcImage)
	case extJpg, extJpeg:
		return jpeg.Encode(dstImage, srcImage, nil)
	case extGif:
		return gif.Encode(dstImage, srcImage, nil)
	default:
		return fmt.Errorf("Failed to encode due to unsupported extension: %v", conv.dstExt)
	}
}

func getSrcImage(srcFilePath string) (image.Image, error) {
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return nil, err
	}
	// TODO Close失敗時の error の返却方法
	defer srcFile.Close()
	srcImage, _, err := image.Decode(srcFile)
	if err != nil {
		return nil, err
	}
	return srcImage, nil
}

func getDstImage(dstFilePath string) (*os.File, error) {
	return os.Create(dstFilePath)
}

func containsStringInSlice(s []string, target string) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

func getType(i interface{}) string {
	return reflect.TypeOf(i).String()
}
