package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytesRead, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", bytesRead)
}
