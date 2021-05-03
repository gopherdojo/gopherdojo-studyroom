package omikuji

import (
	"fmt"
	"math/rand"
	"time"
)

var result = []string{"大凶", "凶", "吉", "中吉", "大吉"}

func Run() {
	fmt.Println("start omikuji")
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(result))
	r := result[i]
	fmt.Println(r)
}
