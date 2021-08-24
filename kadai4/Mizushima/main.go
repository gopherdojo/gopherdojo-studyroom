package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

type Record struct {
	ID int64
	Name string
	Phone string
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	db, err := sql.Open("sqlite", "addressbook.db")
	if err != nil {
		return err
	}


	if err = db.Ping(); err != nil {
		return err
	} else {
		fmt.Fprintln(os.Stderr, "connected to dakabase.db")
	}

	if err := createTable(db); err != nil {
		return err
	}

	for {
		if err := showRecords(db); err != nil {
			return err
		}

		if err := inputRecord(db); err != nil {
			return err
		}
	}

	return nil
}

func createTable(db *sql.DB) error {
	const sql = `
	CREATE TABLE IF NOT EXISTS addressbook (
		id   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		phone  INTEGER NOT NULL
	);`

	if _, err := db.Exec(sql); err != nil {
		return err
	}

	return nil
}

func showRecords(db *sql.DB) error {
	fmt.Println("All records.")
	rows, err := db.Query("SELECT * FROM addressbook")
	if err != nil {
		return err
	}

	for rows.Next() {
		var r Record
		if err := rows.Scan(&r.ID, &r.Name, &r.Phone); err != nil {
			return err
		}
		fmt.Println(r)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func inputRecord(db *sql.DB) error {
	var r Record

	fmt.Print("Name? > ")
	fmt.Scan(&r.Name)

	fmt.Print("Phone number? > ")
	fmt.Scan(&r.Phone)

	const sql = "INSERT INTO addressbook(name, phone) values (?,?)"
	_, err := db.Exec(sql, r.Name, r.Phone)
	if err != nil {
		return err
	}

	return err
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
// 	BootServer(handler)

// }

// func BootServer(handler func(http.ResponseWriter, *http.Request)) {

// 	http.HandleFunc("/", handler)
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		fmt.Printf("BootServer: err: %s\n", err)
// 	}

// }
