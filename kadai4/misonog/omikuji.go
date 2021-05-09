package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	timeForm = "2006-1-2"
	daikichi = "大吉"
)

var omikujiResults = []string{"吉", "小吉", "凶"}

type omikujiResult struct {
	Result string `json:"result"`
}

type omikuji struct {
	result string
	date   time.Time
}

func handler(w http.ResponseWriter, r *http.Request) {
	var o omikuji

	dateValue := r.FormValue("date")
	if dateValue != "" {
		date, err := time.Parse(timeForm, dateValue)
		if err != nil {
			log.Fatal(err)
		}
		o = omikuji{date: date}
	} else {
		o = omikuji{}
	}

	o.Draw()
	data := &omikujiResult{Result: o.result}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func (o *omikuji) Draw() {
	if isShogatsu(o.date) {
		o.result = daikichi
	} else {
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(len(omikujiResults))
		o.result = omikujiResults[i]
	}
}

func isShogatsu(t time.Time) bool {
	// 0001/1/1は正月と判定できない
	if t.IsZero() {
		return false
	}

	_, month, date := t.Date()
	if month == time.January {
		if date == 1 || date == 2 || date == 3 {
			return true
		}
	}
	return false
}
