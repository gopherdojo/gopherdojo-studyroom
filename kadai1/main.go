package main

import "os"

func main() {
	cli := &CLI{inStream: os.Stdin, outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
