package imgconv

import (
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
func (i ImageConverter) Exec() error {
	dirPaths, err := collectDirPaths(i.path)
	if err != nil {
		return err
	}

	fileContainers := makeFileContainers(dirPaths)
	filteredFileContainers := filterFileContainers(fileContainers, i.from)
	convert(filteredFileContainers, i.dist, i.from, i.to)

	fmt.Println("the image conversion to " + i.from.string() + " + => " + i.to.string() + " was successful!!")

	return nil
}

// collectDirPaths 指定されたパス配下にあるディレクトリを再帰的に取得する
func collectDirPaths(path string) ([]string, error) {
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
		return nil, err
	}
	return dirPaths, nil
}

// collectFilesOfDir ディレクトリ配下のファイルを取得する
func collectFilesOfDir(path string) ([]string, error) {
	filesOfDir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.New("could not read " + path)
	}

	var files []string
	for _, file := range filesOfDir {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}
	return files, nil
}

// fileContainer ディレクトリパスと配下のファイルを管理する
type fileContainer struct {
	dirPath   string
	filesName []string
}

// makeFileContainers ディレクトとファイルの連想配列を生成する
func makeFileContainers(dirPaths []string) []fileContainer {
	fileContainers := []fileContainer{}
	fmt.Println("##### list of files that failed to read #####")
	for _, dirPath := range dirPaths {
		files, err := collectFilesOfDir(dirPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		fileContainer := fileContainer{dirPath, files}
		fileContainers = append(fileContainers, fileContainer)
	}
	fmt.Print("\n\n")
	return fileContainers
}

// filterFileContainers fromで指定されたフォーマットに絞り込む
func filterFileContainers(targets []fileContainer, from Format) []fileContainer {
	filterdFileContainer := []fileContainer{}
	fmt.Println("##### list of files that could not be opened or could not be decode config #####")
	for _, t := range targets {
		files := []string{}
		for _, v := range t.filesName {
			file, err := os.Open(filepath.Join(t.dirPath, v))
			if err != nil {
				fmt.Fprintln(os.Stderr, v+": reason: could not open the file")
				continue
			}

			if _, err := decodeConfig(file, from); err != nil {
				fmt.Fprintln(os.Stderr, v+": reason could not be decode config")
				continue
			}
			files = append(files, v)
		}
		fileContainer := fileContainer{t.dirPath, files}
		filterdFileContainer = append(filterdFileContainer, fileContainer)
	}
	fmt.Print("\n\n")
	return filterdFileContainer
}

// decodeConfig エンコードされた画像のカラーモデルと寸法をデコードする
func decodeConfig(r io.Reader, from Format) (image.Config, error) {
	switch from {
	case _jpg:
		return jpeg.DecodeConfig(r)
	case _png:
		return png.DecodeConfig(r)
	case _gif:
		return gif.DecodeConfig(r)
	default:
		return image.Config{}, errors.New("decode config failed")
	}
}

// convert 画像形式を変換する
func convert(targets []fileContainer, dist string, from Format, to Format) {
	fmt.Println("##### list of files that failed to convert #####")
	for _, t := range targets {
		for _, fn := range t.filesName {
			img, err := decode(filepath.Join(t.dirPath, fn), from)
			if err != nil {
				fmt.Fprint(os.Stderr, fn+": reason: decoding failed")
				continue
			}

			if err := encode(img, dist, fn, from, to); err != nil {
				fmt.Fprint(os.Stderr, fn+": reason: encoding failed")
				continue
			}
		}
	}
	fmt.Print("\n\n")
}

// decode ファイルを画像オブジェクトに変換する
func decode(path string, from Format) (image.Image, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
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
func GetFormatEnum(ext string) Format {
	switch ext {
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
