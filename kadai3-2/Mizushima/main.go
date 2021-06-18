package main

import (
	"fmt"
	"log"
	"os"

	"github/MizushimaToshihiko/gopherdojo-studyroom/kadai3-2/Mizushima/options"
)

// go run main.go -url https://4.bp.blogspot.com/-2-Ny23XgrF0/Ws69gszw2jI/AAAAAAABLdU/unbzWD_U8foWBwPKWQdGP1vEDoQoYjgZwCLcBGAs/s1600/top_banner.jpg


func main() {
	opts, err := options.ParseArgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(opts)

	// fmt.Println("url:", *url)
	// *filePath += filepath.Base(*url)
	// // fmt.Println("file_path:", *file_path)

	// if err := DownloadFile(*filePath, *url); err != nil {
	// 	log.Fatal(err)
	// }
}

// func DownloadFile(filepath string, url string) error {
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return errors.Wrap(err, fmt.Sprintf("failed to split NewRequest for get: %d", r.worker))
// 	}

// 	fmt.Printf("%#v\n", req)

// 	req.Header.Set("Range", fmt.Sprintf())

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	out, err := os.Create(filepath)
// 	if err != nil {
// 		return err
// 	}
// 	defer out.Close()

// 	_, err = io.Copy(out, resp.Body)
// 	return err
// }