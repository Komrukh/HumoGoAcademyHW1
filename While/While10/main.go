package main

import (
	"fmt"
)

func main() {
	var n int
	k := 1
	a := 3
	fmt.Scan(&n)
	for a < n {
		a *= 3
		k++
	}
	fmt.Println(k - 1)
}
