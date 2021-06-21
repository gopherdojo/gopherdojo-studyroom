package download

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(path string, urls []string) error {

	for _, url := range urls {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		// fmt.Printf("%#v\n", req)

		req.Header.Set("Range", "byte=0-499")

		resp, err := http.DefaultClient.Do(req)
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