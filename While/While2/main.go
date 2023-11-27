package main

import "fmt"

func main() {
	var k, m, u float64
	fmt.Scan(&k, &m)
	for k > m {
		k -= m
		u++
	}
	fmt.Println(u)
}
