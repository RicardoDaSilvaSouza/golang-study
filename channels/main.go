package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
		"http://www.amazon.com",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}

	/* for {
		go checkLink(<-channel, channel)
	} */

	for link := range channel {
		go checkLink(link, channel)
	}
}

func checkLink(link string, channel chan string) {
	time.Sleep(time.Second * 5)
	if _, err := http.Get(link); err != nil {
		fmt.Println(link, "might be down!")
		channel <- link
	} else {
		fmt.Println(link, "is ok!")
		channel <- link
	}
}
