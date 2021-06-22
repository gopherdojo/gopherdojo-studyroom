package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":12345", nil)
}