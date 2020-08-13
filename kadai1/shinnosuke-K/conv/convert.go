package conv

import (
	"fmt"
)

func Do(dirPath string, before string, after string) {
	files := getFiles(dirPath)
	for n := range files {
		fmt.Println(files[n])
	}
	fmt.Println(before, after)
}
