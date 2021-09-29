package main

import k1 "github.com/karasuex54/kadai1"

func main() {
	dirPath := "/home/ubuntu/MyWork/gopherdojo-studyroom"
	c := k1.NewConverter(dirPath, k1.ToExt("gifa"), k1.FromExt("png"))
	c.Run()
}
