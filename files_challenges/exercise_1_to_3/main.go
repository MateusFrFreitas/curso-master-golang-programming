package main

import (
	"log"
	"os"
)

func main() {
	// 1.
	file, err := os.Create("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// 2.
	_, err = os.Stat(file.Name())
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist!")
		}
	}

	if err := os.Rename(file.Name(), "data.txt"); err != nil {
		log.Fatal(err)
	}

	// 3.
	if err := os.Remove("data.txt"); err != nil {
		log.Fatal(err)
	}
}
