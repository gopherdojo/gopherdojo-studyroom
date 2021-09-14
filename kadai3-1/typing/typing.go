package typing

import (
    "math/rand"
    "time"
    "os"
    "bufio"
)

func RandomWord() string {
        words := []string {
            "reason",
            "secret",
            "gimlet",
            "escape",
            "galaxy",
            "breeze",
            "beetle",
            "allure",
            "velvet",
        }

        rand.Seed(time.Now().UnixNano())

        return words[rand.Intn(len(words))]
}

func CreateChan(word string) <-chan string {
        stdin := bufio.NewScanner(os.Stdin)
        stdin.Scan()
        text := stdin.Text()
        c := make(chan string, 1)
        c <- text
        return c
}
