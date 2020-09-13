package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 標準入力から受け取ったテキストに行番号をつける
	line := 0
	for scanner.Scan() {
		line += 1
		fmt.Fprintf(os.Stdout, "%6d: %s\n", line, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "read error :%v", err)
	}
}
