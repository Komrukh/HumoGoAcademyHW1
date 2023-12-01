package main

import (
	"fmt"
)

func main() {
	var n, k, l int
	fmt.Scan(&n, &k)
	for n <= k {
		n = n * 3
		k = k * 2
		l++
	}
	fmt.Println(l)
}
