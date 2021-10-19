package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, c chan string) {
	resp, err := http.Get(url)

	var s string

	if err != nil {
		s = fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error: %v\n", err)
		c <- s
		return
	}

	defer resp.Body.Close()
	s = fmt.Sprintf("%s -> Status Code: %d\n", url, resp.StatusCode)

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			s += "Error reading the body\n"
			c <- s
			return
		}

		file := strings.Split(url, "//")[1]
		file += ".txt"

		s += fmt.Sprintf("Wrting response body to %s\n", file)

		err = ioutil.WriteFile(file, bodyBytes, 0664)
		if err != nil {
			s += "Error writting file\n"
			c <- s
			return
		}

		s += fmt.Sprintf("%s is UP\n", url)

	}

	c <- s
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://www.google.com",
		"https://www.medium.com",
		"https://www.golang1.com",
		"https://www.google.com/a.html",
	}

	ch := make(chan string)
	defer close(ch)

	for _, url := range urls {
		go checkAndSaveBody(url, ch)
	}

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-ch)
	}
}
