package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleFnLate(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Accept-Range", "bytes")
	time.Sleep(30*time.Second)
	fmt.Fprint(w, "Hello\n")
}

func handleFn(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Accept-Range", "bytes")
	// time.Sleep(30*time.Second)
	fmt.Fprint(w, "Hello\n")
}

func main() {
	http.HandleFunc("/late", handleFnLate)
	http.HandleFunc("/", handleFn)
	http.ListenAndServe(":12345", nil)
}