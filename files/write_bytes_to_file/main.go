package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	byteSlice := []byte("I learn Golang!")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bytes written: %d\n", bytesWritten)

	bs := []byte("Go Programming is cool!")
	err = ioutil.WriteFile("c.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
