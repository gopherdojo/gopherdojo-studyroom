package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kotaaaa/gopherdojo-studyroom/kadai4/kotaaaa/fortune"
)

// API response model
type ResModel struct {
	Status string `json:"status"`
	Result string `json:"result"`
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	p := &ResModel{Status: "", Result: ""}
	var t time.Time
	var err error
	date := r.FormValue("p")
	timeFormat := "2006-01-02"

	// If request parameter exists
	fmt.Println("date:", date)
	if date != "" {
		t, err = time.Parse(timeFormat, date)
		if err != nil {
			fmt.Println(fmt.Errorf("Time format invalid: %v", err.Error()))
			// w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		t = time.Now()
	}

	res := fortune.Draw(t)
	p.Status = "Success"
	p.Result = res.String()
	var buf bytes.Buffer
	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewEncoder(&buf).Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

//Run application
func Run() {
	http.HandleFunc("/draw", httpHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
