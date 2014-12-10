package main

import "fmt"

type customSlice []string

func wantSlice(s customSlice) {
	fmt.Println(s)
}

func main() {
	slice := []string{`a`, `b`}
	wantSlice(slice)
	// [a b]
}
