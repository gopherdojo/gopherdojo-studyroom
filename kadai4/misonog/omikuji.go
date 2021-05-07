package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type omikujiResult struct {
	Result string `json:"result"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := &omikujiResult{Result: "吉"}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}
