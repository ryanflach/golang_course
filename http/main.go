package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	r, e := http.Get("http://www.google.com")

	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(1)
	}

	io.Copy(os.Stdout, r.Body)
}
