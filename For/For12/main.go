package main

import (
	"fmt"
)

func main() {
	var n int
	var s, i float32
	fmt.Scan(&n)
	s = 1
	for i = 1.1; i <= float32(n)/10+1.1; i += 0.1 {
		s *= i
	}
	fmt.Println(s)
}
