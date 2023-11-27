package main

import (
	"fmt"
)

func main() {
	var n, s, i int
	fmt.Scan(&n)
	for i = 1; i <= 2*n-1; i += 2 {
		s += i
		fmt.Println(s)
	}
}
