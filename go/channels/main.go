package main

import (
	"fmt"
	"net/http"
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

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!!!")
		c <- "ST down"
		return
	}

	fmt.Println(link, "is up!")
	c <- "ST up"
}
