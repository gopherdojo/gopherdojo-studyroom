package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var res = map[int]string{
	0: "大吉",
	1: "中吉",
	2: "中吉",
	3: "小吉",
	4: "小吉",
	5: "凶",
}

type Person struct {
	Name string `json:"name"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")

	t := template.Must(template.New("msg").
		Parse("<!DOCTYPE html><html><body>{{.Person}}さんの運勢は「<b>{{.Response}}</b>」です</body></html>"))

	if err := t.Execute(w, struct {
		Person   string
		Response string
	}{
		Person:   r.FormValue("p"),
		Response: res[rand.Intn(len(res))],
	}); err != nil {
		log.Fatalf("failed to execute template: %v", err)
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
