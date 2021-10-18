package main

import (
	"io/ioutil"
	"log"
)

func main() {
	err := ioutil.WriteFile("info.txt", []byte("The Go gopher is an iconic mascot!"), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
