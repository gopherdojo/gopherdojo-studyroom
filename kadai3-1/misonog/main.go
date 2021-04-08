package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
