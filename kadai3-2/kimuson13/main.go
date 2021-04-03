package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/kimuson13/gopherdojo-studyroom/kimuson13/download"
)

func main() {
	size, err := download.GetContentLength("https://www.naoshima.net/wp-content/uploads/2020/06/786619bb442b57802bccc419e9d2e381.pdf")
	if err != nil {
		fmt.Fprintf(os.Stderr, "something happen: %v", err)
		os.Exit(1)
	}
	fmt.Printf("size: %v\n", size)
	num := runtime.NumCPU()
	fmt.Printf("NumCPU: %v\n", num)
	parallel := size / num
	fmt.Printf("Parallel: %v", parallel)
}
