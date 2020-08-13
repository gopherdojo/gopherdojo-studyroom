package conv

import (
	"fmt"
)

func Do(dirPath string, before string, after string) {
	getFiles(dirPath)
	fmt.Println(before, after)
}
