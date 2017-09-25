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
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(l string) {
	_, e := http.Get(l)

	if e != nil {
		fmt.Println(l, "might be down:", e)
		return
	}

	fmt.Println(l, "is up and running!")
}
