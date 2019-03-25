package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://goland.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go chechLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func chechLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, " is up")
	c <- "Yep is up"
}