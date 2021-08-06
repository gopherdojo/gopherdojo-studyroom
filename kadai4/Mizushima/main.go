package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var res = map[int]string{
	0: "大吉",
	1: "中吉",
	2: "小吉",
	3: "凶",
}

type Person struct {
	Name string `json:"name"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")

	person := r.FormValue("p")

	idx := rand.Intn(len(res))

	v := struct {
		Msg string `json:"msg"`
	}{
		Msg: fmt.Sprintf("%sさんの運勢は%sです", person, res[idx]),
	}

	// JSONを返す（レスポンスに書き込む）
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Err:", err)
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
