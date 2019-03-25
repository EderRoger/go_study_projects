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

	for _, link := range links {
		chechLink(link)
	}
}

func chechLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		return
	}

	fmt.Println(link, " is up")
}
