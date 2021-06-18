/*This is a training code for learning Go language.
And this code can change the format of images as you like*/
package main

import (
	"bufio"
	"convert"
	"fmt"
	"os"
	"strings"
)

type oriPath struct {
	inputPA  string
	outputPA string
}

func main() {
	var a oriPath
	a.inputPA = "./img"
	a.outputPA = "./img"
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input the image format you want(like png or gif):")
	if !scanner.Scan() {
		fmt.Println("Please input the format.")
		return
	}
	srcPath := scanner.Text()

	var s []string
	s, _ = convert.GetAllFile(a.inputPA, s)
	for _, i := range s {
		j := strings.Replace(i, "jpg", srcPath, -1) //jpgから任意の形式へ変換する
		err := convert.Conv(i, j)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		}
	}
}
