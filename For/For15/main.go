package main

import (
	"fmt"
)

func main() {
	var n, i int
	var s float64
	k := 1.0
	fmt.Scan(&s, &n)
	for i = 0; i < n; i++ {
		k *= s
	}
	fmt.Println(k)
}
