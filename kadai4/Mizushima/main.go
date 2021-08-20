package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintln(os.Stderr, "connected to dakabase.db")
	}
	
}

// import (
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"time"
// )

// var res = map[int]string{
// 	0: "大吉",
// 	1: "中吉",
// 	2: "中吉",
// 	3: "小吉",
// 	4: "小吉",
// 	5: "凶",
// }

// type Person struct {
// 	Name string `json:"name"`
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "text/html; charset=utf-8")

// 	t := template.Must(template.New("msg").
// 		Parse("<!DOCTYPE html><html><body>{{.Person}}さんの運勢は「<b>{{.Response}}</b>」です</body></html>"))

// 	if err := t.Execute(w, struct {
// 		Person   string
// 		Response string
// 	}{
// 		Person:   r.FormValue("p"),
// 		Response: res[rand.Intn(len(res))],
// 	}); err != nil {
// 		log.Fatalf("failed to execute template: %v", err)
// 	}

// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	BootServer(handler)

// }

// func BootServer(handler func(http.ResponseWriter, *http.Request)) {

// 	http.HandleFunc("/", handler)
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		fmt.Printf("BootServer: err: %s\n", err)
// 	}

// }
