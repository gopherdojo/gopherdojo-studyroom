package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

const (
	// Extensions
	ExtJpg        = ".jpg"
	ExtJpeg       = ".jpeg"
	ExtPng        = ".png"
	ExtGif        = ".gif"
	DefaultSrcExt = ExtJpg
	DefaultDstExt = ExtPng
	// MIME Types
	MimeTypeJpeg = "image/jpeg"
	MimeTypePng  = "image/png"
	MimeTypeGif  = "image/gif"
	// Directories
	DefaultSrcDir = "./testdata/src"
	DefaultDstDir = "./testdata/dst"
	// Flag
	DefaultFileDeleteFlag = false
)

type Converter struct {
	SrcExt         string
	DstExt         string
	SrcDir         string
	DstDir         string
	FileDeleteFlag bool
}

// Mapping table of extensions and MIME types
/*
var ExtMimeTypeList = map[string]string{
	ExtJpg:  MimeTypeJpeg,
	ExtJpeg: MimeTypeJpeg,
	ExtPng:  MimeTypePng,
	ExtGif:  MimeTypeGif,
}
*/

var (
	// Extensions' list
	ExtList    []string = []string{ExtJpg, ExtJpeg, ExtPng, ExtGif}
	ExtListStr string   = strings.Join(ExtList, " ")
	// Arguments
	SrcExt         *string = flag.String("src-ext", DefaultSrcExt, "Source extension (choices "+ExtListStr+")")
	DstExt         *string = flag.String("dst-ext", DefaultDstExt, "Destination extension (choices "+ExtListStr+")")
	SrcDir         *string = flag.String("src-dir", DefaultSrcDir, "Source directory")
	DstDir         *string = flag.String("dst-dir", DefaultDstDir, "Destination directory")
	FileDeleteFlag *bool   = flag.Bool("delete", DefaultFileDeleteFlag, "File delete flag")
)

func main() {
	flag.Parse()
	if err := ValidArgs(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	conv := NewConverter(*SrcExt, *DstExt, *SrcDir, *DstDir, *FileDeleteFlag)
	err := conv.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func ValidArgs() error {
	if !ContainsStringInSlice(ExtList, *SrcExt) {
		return fmt.Errorf("The selected \"src-ext\" does not exist in: %v", ExtListStr)
	}

	if !ContainsStringInSlice(ExtList, *DstExt) {
		return fmt.Errorf("The selected \"dst-ext\" does not exist in: %v", ExtListStr)
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

	if GetType(*FileDeleteFlag) != "bool" {
		return errors.New("The \"delete\" option must be true or false")
	}
	return nil
}

func NewConverter(srcExt string, dstExt string, srcDir string, dstDir string, fileDeleteFlag bool) *Converter {
	return &Converter{
		SrcExt:         srcExt,
		DstExt:         dstExt,
		SrcDir:         srcDir,
		DstDir:         dstDir,
		FileDeleteFlag: fileDeleteFlag,
	}
}

func (conv *Converter) Run() error {
	return filepath.Walk(conv.SrcDir, func(srcFilePath string, info os.FileInfo, err error) error {
		// filepath.Walk() 内部で os.FileInfo 取得時にエラーがある場合、処理をスキップ
		if err != nil {
			return err
		}
		// ディレクトリの場合、処理をスキップ
		if info.IsDir() {
			return nil
		}
		// 取得した拡張子と変換元の拡張子が異なる場合、処理をスキップ
		if ext := GetFormattedFileExt(srcFilePath); ext != conv.SrcExt {
			return nil
		}
		// 変換後のディレクトリ・パスを取得
		dstDir, err := conv.GetDstDir(srcFilePath)
		if err != nil {
			return err
		}
		// 変換後のディレクトリが存在しない場合は作成 (os.MkdirAll() 内でディレクトリの有無を判定)
		if err = os.MkdirAll(dstDir, 0777); err != nil {
			return err
		}
		// 変換後のファイル・パスを取得
		dstFilePath := conv.GetDstFilePath(srcFilePath, dstDir)
		// 画像の変換処理
		if err := conv.Convert(srcFilePath, dstFilePath); err != nil {
			return err
		}
		// 変換元の画像を削除
		if conv.FileDeleteFlag {
			if err := DeleteSrcFile(srcFilePath); err != nil {
				return err
			}
		}
		return nil
	})
}

func DeleteSrcFile(srcFilePath string) error {
	return os.Remove(srcFilePath)
}

func (conv *Converter) GetDstDir(srcFilePath string) (string, error) {
	// 変換元の対象ディレクトリの相対パスを取得
	srcRelDir, err := conv.GetSrcRelDir(srcFilePath)
	if err != nil {
		return "", err
	}
	// 上記の相対パスと変換後の基準ディレクトリを元に
	// 変換後のディレクトリの絶対パスを取得
	path := filepath.Join(conv.DstDir, srcRelDir)
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return absPath, nil
}

func (conv *Converter) GetSrcRelDir(srcFilePath string) (string, error) {
	dir := GetDirName(srcFilePath)
	return GetRelPath(conv.SrcDir, dir)
}

func (conv *Converter) GetDstFilePath(srcFilePath string, dstDir string) string {
	dstFileName := GetFileStem(srcFilePath) + conv.DstExt
	return filepath.Join(dstDir, dstFileName)
}

func (conv *Converter) Convert(srcFilePath string, dstFilePath string) error {
	srcImage, err := GetSrcImage(srcFilePath)
	if err != nil {
		return err
	}
	dstImage, err := GetDstImage(dstFilePath)
	if err != nil {
		return err
	}
	if err := conv.Encode(dstImage, srcImage); err != nil {
		return err
	}
	return nil
}

func (conv *Converter) Encode(dstImage io.Writer, srcImage image.Image) error {
	switch conv.DstExt {
	case ExtPng:
		return png.Encode(dstImage, srcImage)
	case ExtJpg, ExtJpeg:
		return jpeg.Encode(dstImage, srcImage, nil)
	case ExtGif:
		return gif.Encode(dstImage, srcImage, nil)
	default:
		return fmt.Errorf("Failed to encode due to unsupported extension: %v", conv.DstExt)
	}
}

func GetSrcImage(srcFilePath string) (image.Image, error) {
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

func GetDstImage(dstFilePath string) (*os.File, error) {
	return os.Create(dstFilePath)
}

/*
以下、汎用性の高いメソッド
*/
func ContainsStringInSlice(s []string, target string) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

func GetType(i interface{}) string {
	return reflect.TypeOf(i).String()
}

func GetDirName(path string) string {
	return filepath.Dir(path)
}

func GetRelPath(basePath string, targetPath string) (string, error) {
	relPath, err := filepath.Rel(basePath, targetPath)
	if err != nil {
		return "", err
	}
	return relPath, nil
}

func GetFileName(path string) string {
	return filepath.Base(path)
}

func GetFileStem(path string) string {
	pathLength := len(path)
	extLength := len(filepath.Ext(path))
	return filepath.Base(path[:pathLength-extLength])
}

func GetFormattedFileExt(path string) string {
	ext := GetFileExt(path)
	return FormatExt(ext)
}

func GetFileExt(path string) string {
	return filepath.Ext(path)
}

func FormatExt(ext string) string {
	return strings.ToLower(ext)
}

func GetMimeType(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	buf := make([]byte, 512)
	if _, err := file.Read(buf); err != nil {
		return "", err
	}
	mimeType := http.DetectContentType(buf)
	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}
	return mimeType, nil
}
