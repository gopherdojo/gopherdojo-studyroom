package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sorakoro/gopherdojo-studyroom/kadai1/imgconv"
	// "github.com/sorakoro/gopherdojo-studyroom/kadai1/imgconv"
)

var (
	path = flag.String("p", "", "relative path to the image to convert")
	dist = flag.String("d", "", "path to save the converted image")
	from = flag.String("f", "jpg", "extension before conversion. The default value is jpg")
	to   = flag.String("t", "png", "extension after conversion. The default value is png")
)

var supportFormats = [3]string{"jpg", "png", "gif"}

func main() {
	flag.Parse()

	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if _, err := checkFormat(strings.ToLower(*from)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if _, err := checkFormat(strings.ToLower(*to)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	converter := imgconv.NewImageConverter(
		filepath.Join(wd, *path),
		filepath.Join(wd, *dist),
		imgconv.GetFormatEnum(*from),
		imgconv.GetFormatEnum(*to),
	)
	converter.Exec()
}

// checkFormat 引数で渡された画像フォーマットをサポートしているかチェック
func checkFormat(format string) (bool, error) {
	for _, sf := range supportFormats {
		if format == sf {
			return true, nil
		}
	}
	return false, errors.New("not a supported format " + format)
}
