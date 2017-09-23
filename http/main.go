package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	r, e := http.Get("http://www.google.com")

	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(1)
	}
	lw := logWriter{}

	io.Copy(lw, r.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
