package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkUrl(url string, c chan string) {
	defer func() {
		c <- url
	}()

	resp, err := http.Get(url)

	var s string

	if err != nil {
		s = fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error: %v\n", err)
		fmt.Println(s)
		return
	}

	defer resp.Body.Close()
	s = fmt.Sprintf("%s -> Status Code: %d\n", url, resp.StatusCode)
	fmt.Println(s)
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
		go checkUrl(url, ch)
	}

	// for {
	// 	go checkUrl(<-ch, ch)
	// 	fmt.Println(strings.Repeat("#", 30))
	// 	time.Sleep(2 * time.Second)
	// }

	// for url := range ch {
	// 	time.Sleep(2 * time.Second)
	// 	go checkUrl(url, ch)
	// }

	for url := range ch {
		go func(u string) {
			time.Sleep(2 * time.Second)
			checkUrl(u, ch)
		}(url)
	}
}
