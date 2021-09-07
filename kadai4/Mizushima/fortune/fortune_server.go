// fortune_server.go implements the 'handler' and other functions to return a result of 'Omikuji'.
package fortune

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var timeNow = func() time.Time { return time.Now() }
var resOmikuji = [4]string{"大吉", "中吉", "小吉", "凶"}

// Res is a struct for json has one field 'result'.
type Res struct {
	Result string `json:"result"`
}

// omikujiHandler implements the 'handler' for the 'Omikuji' server.
func OmikujiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")

	v := Res{
		Result: result(rand.Intn(5), timeNow()),
	}

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Error: ", err)
	}

}

// result function returns the result of Omikuji if 'i' is bitween 0 and 5, or a empty string.
// If 't' is the first three days of the new year, result function returns always '大吉'
func result(i int, t time.Time) string {

	if t.Month() == 1 && (1 <= t.Day() && t.Day() <= 3) {
		return resOmikuji[0]
	}

	switch i {
	case 0:
		return resOmikuji[0]
	case 1, 2:
		return resOmikuji[1]
	case 3, 4:
		return resOmikuji[2]
	case 5:
		return resOmikuji[3]
	default:
		return ""
	}
}
