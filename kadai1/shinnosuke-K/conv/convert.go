package conv

import (
	"fmt"
	"os"
	"strings"
)

var imgExts = []string{"gif", "png", "jpeg"}

func Do(dirPath string, before string, after string) {
	files := getFiles(dirPath)
	for n := range files {
		fmt.Println(files[n])
	}

	err := checkOpt(before, after)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkOpt(before string, after string) error {
	for n := range imgExts {
		if strings.ToLower(before) != imgExts[n] || strings.ToLower(after) != imgExts[n] {
			return fmt.Errorf("imgconv: invaild image extension")
		}
	}
	return nil
}
