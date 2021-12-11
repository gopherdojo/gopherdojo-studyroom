package typing

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func RandomWord() string {
	words := []string {"archive", "bufio", "builtin", "bytes", "cmd",
		"compress", "container", "context", "crypto", "database", "debug", "embed",
		"encoding", "errors", "expvar", "flag", "fmt", "go", "hash", "html", "image",
		"index", "io", "log", "math", "mine", "net", "os", "path", "plugin", "reflect",
		"regexp", "runtime", "sort", "strconv", "strings", "sync", "syscall",
		"testing", "text", "time", "unicode", "unsafe"}

	rand.Seed(time.Now().UnixNano())

	return words[rand.Intn(len(words))]
}

func CreateChan() <-chan string {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	channel := make(chan string, 1)
	channel <- text
	return channel
}
