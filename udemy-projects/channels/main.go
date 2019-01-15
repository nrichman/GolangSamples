package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(c, link)
	}

	// Wait for the channel to return some value, after it does - assign it to l
	for l := range c {
		go func(link string) { // l can change without the go routine referencing the same as main
			time.Sleep(5 * time.Second)
			checkLink(c, link)
		}(l) // <-- call the function after declaring it
	}
}

func checkLink(c chan string, link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
