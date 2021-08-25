package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	BootServer(omikujiHandler)
}

func BootServer(handler func(http.ResponseWriter, *http.Request)) {

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("BootServer: err: %s\n", err)
	}

}
