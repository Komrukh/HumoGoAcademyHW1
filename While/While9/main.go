package main

import (
	"fmt"
)

func main() {
	var k int
	a := 1
	n := 3
	fmt.Scan(&k)
	for n <= k {
		n *= 3
		a++
	}
	fmt.Println(a)
}
