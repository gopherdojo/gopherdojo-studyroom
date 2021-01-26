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

var results = []string{"凶", "吉", "中吉", "大吉"}

type Omikuji struct {
	Result string `json:"result"`
}

func pickResult() string {
	_, m, d := time.Now().Date()
	var i int
	if m == 1 && d > 0 && d < 4 {
		i = len(results) - 1
	} else {
		i = rand.Intn(len(results))
	}
	return results[i]
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	result := pickResult()
	o := &Omikuji{Result: result}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(o); err != nil {
		log.Fatal(err)
	}
	str := buf.String()
	log.Println(str)
	fmt.Fprintf(w, "%s", str)
}

func Run() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", handler)
	log.Printf("omikuji api is listiing port 8080...\n")
	http.ListenAndServe(":8080", nil)
}
