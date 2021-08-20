package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	response, err := http.Get("http://localhost:8080?p=Gopher")
	if err != nil {
		log.Fatalf("main: err: %s", err)
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
