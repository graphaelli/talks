package main

import "fmt"

type customInt int

func wantInt(i customInt) {
	fmt.Println(i)
}

func main() {
	integer := 5
	// cannot use integer (type int) as type customInt in argument to wantInt
	// wantInt(integer)
	wantInt(customInt(integer))
}
