package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://www.oracle.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for _, link := range links {
		fmt.Println(<-c, link)
	}

}

func checkLink(link string, c chan string) {
	_, error := http.Get(link)

	if error != nil {
		fmt.Println(link, "might be down")
		c <- "might be down"
		return
	}
	fmt.Println(link, "link is up")
	c <- "link is up"

}
