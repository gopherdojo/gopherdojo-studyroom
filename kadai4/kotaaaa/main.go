package main

import (
	"math/rand"
	"time"

	"github.com/kotaaaa/gopherdojo-studyroom/kadai4/kotaaaa/handler"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	handler.Run()
}
