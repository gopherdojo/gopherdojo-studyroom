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
	Amount string
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	db, err := sql.Open("sqlite", "amountbook.db")
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

		if err := remittanceProcess(db); err != nil {
			return err
		}
	}

	return nil
}

func createTable(db *sql.DB) error {
	const sql = `
	CREATE TABLE IF NOT EXISTS amountbook (
		id   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		amount INTEGER NOT NULL
	);`

	if _, err := db.Exec(sql); err != nil {
		return err
	}

	return nil
}

func showRecords(db *sql.DB) error {
	fmt.Println("All records.")
	rows, err := db.Query("SELECT * FROM amountbook")
	if err != nil {
		return err
	}

	for rows.Next() {
		var r Record
		if err := rows.Scan(&r.ID, &r.Name, &r.Amount); err != nil {
			return err
		}
		fmt.Println(r)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

// input returns names of sender and reciever, and amount sender wants to send.
func input() (string, string, int64) {
	var user1 string
	fmt.Print("From who? Please enter name. > ")
	fmt.Scan(&user1)

	var user2 string
	fmt.Print("To who? Please enter name. > ")
	fmt.Scan(&user2)

	var amount int64
	fmt.Print("How much Amount you want to send? > ")
	fmt.Scan(&amount)

	return user1, user2, amount
}

func remittanceProcess(db *sql.DB) error {

	sender, reciever, amount := input()
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	rowSender := tx.QueryRow("SELECT * FROM USER = ?", sender)
	var recSender Record
	if err := rowSender.Scan(&recSender.ID, &recSender.Name, &recSender.Amount); err != nil {
		if errRoll := tx.Rollback(); errRoll != nil {
			return fmt.Errorf("%s: %s", err, errRoll)
		}
		return err
	}
	const updateSQLSend = "UPDATE amountbook SET amount = amount - ? WHERE ID = ?"
	if _, err = tx.Exec(updateSQLSend, amount, recSender.ID); err != nil {
		tx.Rollback()
		return err
	}

	rowReciever := tx.QueryRow("SELECT * FROM USER = ?", reciever)
	var recReciever Record
	if err := rowReciever.Scan(&recReciever.Name, &recReciever.Amount); err != nil {
		if errRoll := tx.Rollback(); errRoll != nil {
			return fmt.Errorf("%s: %s", err, errRoll)
		}
		return err
	}
	const updateSQLRecieve = "UPDATE amountbook SET amount = amount + ? WHERE ID = ?"
	if _, err = tx.Exec(updateSQLRecieve, amount, recReciever.ID); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
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
