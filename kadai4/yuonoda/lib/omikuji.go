package omikuji

import (
	"log"
	"math/rand"
	"time"
)

var pattens = []string{"凶", "吉", "中吉", "大吉"}

func Run() {
	log.Print("Run")
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(pattens))
	log.Println(i)
	result := pattens[i]
	log.Println(result)
}
