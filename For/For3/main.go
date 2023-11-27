package main

import "fmt"

func main() {
	var n, m int
	var k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	for i := k - 1; i > n; i-- {
		fmt.Println(i)
		m++
	}
	fmt.Println(m)
}
