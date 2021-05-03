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
}

func pickOmikuji() string {
	i := rand.Intn(len(result))
	r := result[i]
	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	result := pickOmikuji()
	omikuji := &Omikuji{Result: result}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(omikuji); err != nil {
		log.Print(err)
		http.Error(w, "error happen while processing", http.StatusInternalServerError)
	}
	str := buf.String()
	fmt.Fprintln(w, str)
}

func Run() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
