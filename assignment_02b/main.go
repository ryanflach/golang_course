package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File must be provided as command line argument.")
		os.Exit(1)
	}

	c, e := ioutil.ReadFile(os.Args[1])

	if e != nil {
		fmt.Println(e)
		os.Exit(2)
	}

	fmt.Println(string(c))
}
