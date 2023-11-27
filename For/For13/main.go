package main

import (
	"fmt"
)

func main() {
	var n int
	var s, l, k, i float32
	fmt.Scan(&n)
	for i = 1.1; i <= float32(n)/10+1.1; i += 0.2 {
		s += i
	}
	for i = 1.2; i <= float32(n)/10+1.1; i += 0.2 {
		l -= i
	}
	k = l + s
	fmt.Println(k)
}
