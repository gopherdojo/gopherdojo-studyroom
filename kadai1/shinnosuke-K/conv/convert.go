package conv

import (
	"fmt"
	"os"
	"strings"
)

var imgExts = []string{"gif", "png", "jpg", "jpeg"}

func Do(dirPath string, before string, after string) {

	err := checkOpt(before, after)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	files := getImgFiles(dirPath, before)
	for n := range files {
		fmt.Println(files[n])

}

func checkOpt(before string, after string) error {
	for n := range imgExts {
		if strings.ToLower(before) == imgExts[n] || strings.ToLower(after) == imgExts[n] {
			return nil
		}
	}
	return fmt.Errorf("imgconv: invaild image extension")
}
