package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func checkAndSaveBody(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is DOWN!\n", url)
		return
	}

	defer resp.Body.Close()
	fmt.Printf("%s -> Status Code: %d\n", url, resp.StatusCode)

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		file := strings.Split(url, "//")[1]
		file += ".txt"

		fmt.Printf("Wrting response body to %s\n", file)

		err = ioutil.WriteFile(file, bodyBytes, 0664)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://www.google.com",
		"https://www.medium.com",
		"https://www.golang1.com",
		"https://www.google.com/a.html",
	}

	for _, url := range urls {
		checkAndSaveBody(url)
		fmt.Println(strings.Repeat("#", 20))
	}
}
