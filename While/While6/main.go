package main

import (
	"fmt"
)

func main() {
	var k, i int
	f := 1.
	fmt.Scan(&k)
	if k%2 == 0 {
		i = 2
	} else {
		i = 1
	}
	for i <= k {
		f *= float64(i)
		i += 2
	}
	fmt.Println(f)
}
