package sample

import "fmt"

func Foo(s string) string {
	return "aaa " + s + " bbb"
}

func PrintFoo(s string) {
	fmt.Println(s)
}
