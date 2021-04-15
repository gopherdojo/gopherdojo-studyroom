package main

import (
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := run(input(os.Stdin), os.Stdout); err != nil {
		log.Fatal(err)
	}
	// if err := run(); err != nil {
	// 	log.Fatal(err)
	// }
}
