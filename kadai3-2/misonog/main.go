package main

import "log"

func main() {
	url := "https://blog.golang.org/gopher/header.jpg"
	if err := download("header.jpg", url); err != nil {
		log.Fatal(err)
	}
}
