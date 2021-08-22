package main

import (
	"fmt"
)
func main() {

	arguments, err := ParseArguments()
	if err != nil {
		fmt.Println(arguments)
	}
	//
	//if arguments.isHelp {
	//	help()
	//	os.Exit(0)
	//}
	//if err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//	os.Exit(1)
	//}
	//
	//c := NewConvert(arguments.args )
	//
	//if err := c.Run(); err != nil {
	//	fmt.Fprintln(os.Stderr, err)
	//	os.Exit(1)
	//})
}
