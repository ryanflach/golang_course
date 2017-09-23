package main

import "net/http"
import "fmt"
import "os"

func main() {
	r, e := http.Get("http://www.google.com")

	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	r.Body.Read(bs)
	fmt.Println(string(bs))
}
