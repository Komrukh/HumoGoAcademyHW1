package main

import "fmt"

func main() {
	var k, m, u float32
	fmt.Scan(&k, &m)
	for k > m {
		k -= m
		u = k
	}
	fmt.Println(u)
}
