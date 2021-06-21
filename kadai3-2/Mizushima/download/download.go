package download

import (
	"io"
	"os"
	"path/filepath"

	"github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/request"
)

func DownloadFile(path string, urls []string) error {

	for _, url := range urls {
		
		resp, err := request.Request("GET", url, "Range", "bytes=281-294")
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if path[len(path)-1] != '/' {
			path += "/"
		}

		out, err := os.Create(path + filepath.Base(url))
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return err
		}
	}
	return nil
}