package omikuji

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var result = []string{"大凶", "凶", "吉", "中吉", "大吉"}

type Omikuji struct {
	Result string `json:"result"`
	Today  string `json:"today"`
}

func PickOmikuji(t time.Time) string {
	var i int
	_, month, date := t.Date()
	if month == time.January {
		if date == 1 || date == 2 || date == 3 {
			i = 4
		}
	} else {
		i = rand.Intn(len(result))
	}
	r := result[i]
	return r
}

var Layout = "2006-01-02"

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	today := time.Now()
	todayStr := today.Format(Layout)
	result := PickOmikuji(today)
	omikuji := &Omikuji{Result: result, Today: todayStr}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(omikuji); err != nil {
		log.Print(err)
		http.Error(w, "error happen while processing", http.StatusInternalServerError)
	}
	str := buf.String()
	_, err := fmt.Fprintln(w, str)
	if err != nil {
		log.Fatal(err)
	}
}

func Run() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
