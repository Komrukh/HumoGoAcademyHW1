package main

import (
	"fmt"
)

func main() {
	var k, i int
	fmt.Scan(&k)
	for i*i < k {
		i++
	}
	fmt.Println(i)
}
