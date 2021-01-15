package imgconv

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// ImageConverter ファイルに関するデータを管理する
type ImageConverter struct {
	path string
	dist string
	from Format
	to   Format
}

// NewImageConverter ImageConverterを生成する
func NewImageConverter(path string, dist string, from Format, to Format) ImageConverter {
	return ImageConverter{path: path, dist: dist, from: from, to: to}
}

// Exec 画像の変換を実行する
func (i ImageConverter) Exec() {
	dirPaths := collectDirPaths(i.path)
	dirAndFileNameMap := makeDirAndFileNameMap(dirPaths)
	filteredMap := filter(dirAndFileNameMap, i.from)

	for dirPath, files := range filteredMap {
		convert(dirPath, files, i.dist, i.from, i.to)
	}
}

// collectDirPaths 指定されたパス配下にあるディレクトリを再帰的に取得する
func collectDirPaths(path string) []string {
	var dirPaths []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.New("There is no directory for " + path)
		}
		if info.IsDir() {
			dirPaths = append(dirPaths, path)
		}
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return dirPaths
}

// collectFilesOfDir ディレクトリ配下のファイルを取得する
func collectFilesOfDir(path string) ([]string, error) {
	filesOfDir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, file := range filesOfDir {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}
	return files, nil
}

// makeDirAndFileNameMaps ディレクトとファイルの連想配列を生成する
func makeDirAndFileNameMap(dirPaths []string) map[string][]string {
	dirAndFileMap := map[string][]string{}
	for _, dirPath := range dirPaths {
		files, err := collectFilesOfDir(dirPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		dirAndFileMap[dirPath] = files
	}
	return dirAndFileMap
}

// filter fromで指定されたフォーマットに絞り込む
func filter(targets map[string][]string, from Format) map[string][]string {
	maps := map[string][]string{}
	for k, val := range targets {
		for _, v := range val {
			file, err := os.Open(filepath.Join(k, v))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			ok := check(file, from)
			if ok {
				maps[k] = append(maps[k], v)
			}
		}
	}
	return maps
}

// 画像形式をチェックする
func check(r io.Reader, from Format) bool {
	switch from {
	case _jpg:
		return isJpg(r)
	case _png:
		return isPng(r)
	case _gif:
		return isGif(r)
	default:
		return false
	}
}

// isJpg Jpgか確認する
func isJpg(r io.Reader) bool {
	magicnum := []byte{255, 216}
	buf := make([]byte, len(magicnum))

	return isEqual(r, magicnum, buf)
}

// isPng Pngか確認する
func isPng(r io.Reader) bool {
	magicnum := []byte{137, 80, 78, 71}
	buf := make([]byte, len(magicnum))

	return isEqual(r, magicnum, buf)
}

// isGif Gifか確認する
func isGif(r io.Reader) bool {
	magicnum := []byte{71, 73, 70, 56}
	buf := make([]byte, len(magicnum))

	return isEqual(r, magicnum, buf)
}

// isEqual マジックナンバーが同じか確認する
func isEqual(r io.Reader, magicnum []byte, buf []byte) bool {
	_, err := io.ReadAtLeast(r, buf, len(buf))
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return false
	}
	return bytes.Equal(magicnum, buf)
}

// convert 画像形式を変換する
func convert(dirPath string, fileNames []string, dist string, from Format, to Format) {
	for _, fn := range fileNames {
		img, err := decode(filepath.Join(dirPath, fn), from)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}

		if err := encode(img, dist, fn, from, to); err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}
	}
}

// decode ファイルを画像オブジェクトに変換する
func decode(path string, from Format) (image.Image, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return nil, err
	}

	switch from {
	case _jpg:
		return jpeg.Decode(file)
	case _png:
		return png.Decode(file)
	case _gif:
		return gif.Decode(file)
	default:
		return nil, errors.New("decoding failed")
	}
}

// encode 画像ファイルを生成する
func encode(img image.Image, dist string, fileName string, from Format, to Format) error {
	file, err := os.Create(strings.Replace(filepath.Join(dist, fileName), from.string(), to.string(), 1))
	defer file.Close()
	if err != nil {
		return err
	}

	switch to {
	case _jpg:
		return jpeg.Encode(file, img, nil)
	case _png:
		return png.Encode(file, img)
	case _gif:
		return gif.Encode(file, img, nil)
	default:
		return errors.New("encoding failed")
	}
}

// Format 画像フォーマット用のエイリアス
type Format int

// 画像フォーマット用の列挙子
const (
	_jpg Format = iota
	_png
	_gif
	unknown
)

// string enumに対応する文字列を返す
func (f Format) string() string {
	switch f {
	case _jpg:
		return "jpg"
	case _png:
		return "png"
	case _gif:
		return "gif"
	default:
		return "unknown"
	}
}

// GetFormatEnum 画像フォーマット用の列挙子を取得する
func GetFormatEnum(with string) Format {
	switch with {
	case "jpg":
		return _jpg
	case "png":
		return _png
	case "gif":
		return _gif
	default:
		return unknown
	}
}
