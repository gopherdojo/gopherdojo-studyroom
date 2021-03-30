package convert

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type FormatType struct {
	Jpg  string `json: 'JPG'`
	Jpeg string `json: 'JPEG'`
	Png  string `json: 'PNG'`
}

var format FormatType

func init() {
	LoadConfig()
}

func LoadConfig() *FormatType {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot open config file")
		os.Exit(1)
	}
	decoder := json.NewDecoder(file)
	format = FormatType{}
	err = decoder.Decode(&format)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot get configuration from file")
		os.Exit(1)
	}
	return &format
}

func logError(err error, msg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func getFileNameWithoutExt(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
