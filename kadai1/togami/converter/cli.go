package main

import (
	"bufio"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io/ioutil"
	"os"
	"path/filepath"
)

var paths []string

func dirwalk(dir string, from *string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			dirwalk(filepath.Join(dir, file.Name()), from)
			continue
		}
		ext := filepath.Ext(file.Name())
		if ext == *from {
			src := filepath.Join(dir,file.Name())
			paths = append(paths, src)
		}
	}
	return paths
}

func readFile(path string, from *string, to *string) error {
	src, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return err
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	distPath := path[:len(path)-len(*from)] + *to

	out, err := os.Create(distPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return error
	}
	defer out.Close()

	png.Encode(out, img)
	return err
}

func handleFile() string {
	fmt.Println("Would you want to delete the original fail? y/n")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	answer := stdin.Text()
	return answer
}

func deleteFile(path string) error{
	os.Remove(path)
	if err != nil {
		return err
	}
}
