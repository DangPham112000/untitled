package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://stackoverflow.com",
		"http://facebook.com",
		"http://google.com",
		"http://amazon.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(4 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!!!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
