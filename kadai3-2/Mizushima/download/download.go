package download

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"

	"github.com/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/request"
)

var mu sync.Mutex

func DownloadFile(url string, out *os.File) error {
		
		resp, err := request.Request("GET", url, "Range", "bytes=281-294")
		if err != nil {
			return err
		}
		defer resp.Body.Close()

	
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return err
		}
	
	return nil
}

func PDownload(url string, out *os.File, fileSize int, idx int, part int, procs int) error {
	mu.Lock()
	fmt.Printf("%d/%d downloading...\n", idx, procs)
	var start, end int
	if idx == 0 {
		start = 0
	} else {
		start = idx*part + 1
	}

	// 最後だったら
	if idx == procs+1 {
		end = fileSize
	} else {
		end = (idx+1) * part
	}

	bytes := fmt.Sprintf("bytes=%d-%d", start, end)
	resp, err := request.Request("GET", url, "Range", bytes)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	length, err := out.Write(body)
	if err != nil {
		return err
	}
	fmt.Printf("%d/%d was written len=%d\n", idx, procs, length)
	mu.Unlock()
	return nil
}