package main

import "fmt"

func main() {
	a := intSlice(10)

	for _, v := range a {
		fmt.Println(v, "is", numType(v))
	}
}

func intSlice(length int) []int {
	a := make([]int, length+1)

	for i := range a {
		a[i] = i
	}

	return a
}

func numType(num int) string {
	if num%2 == 0 {
		return "even"
	}

	return "odd"
}
