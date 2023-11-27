package main

import (
	"fmt"
)

func main() {
	var k int
	fmt.Scan(&k)
	t := 1
	m := 1
	for m < k {
		m *= 2
		if m == k {
			break
		}
		t++
	}
	fmt.Println(t)
}
