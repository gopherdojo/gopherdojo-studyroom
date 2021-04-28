package main

import "log"

func main() {
	cli := New()
	if err := cli.Run(); err != nil {
		log.Fatal(err)
	}
}
