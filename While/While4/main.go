package main

import (
	"fmt"
)

func main() {
	var k int
	fmt.Scan(&k)
	t := false
	m := 1
	for m < k {
		m *= 3
		if m == k {
			t = true
			//fmt.Println(true)
		} else {
			//fmt.Println(false)
			t = false
		}
	}
	fmt.Println(t)
}
