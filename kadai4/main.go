package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Omikuji struct {
	Month  time.Month `json:"month"`
	Day    int        `json:"day"`
	Result string     `json:"data"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/omikuji", omikujiHandler) //1.path,2.function
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func omikujiHandler(w http.ResponseWriter, r *http.Request) {
	var response Omikuji
	kisi := []string{"大吉", "中吉", "小吉", "末吉", "吉", "凶", "大凶"}
	timeObj := time.Now()
	month := timeObj.Month()
	day := timeObj.Day()
	goodday := [3]int{1, 2, 3}
	for _, d := range goodday {
		if month == 1 && day == d {
			response = Omikuji{month, day, kisi[0]}
		} else {
			response = Omikuji{month, day, kisi[rand.Intn(7)]}
		}
	}

	res, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res)) //デバイスファイルに出力
}
