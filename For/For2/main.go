package main

import "fmt"

func main() {
	var n, m int
	var k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	for i := n; i <= k; i++ {
		fmt.Println(i)
		m++
	}
	fmt.Println(m)
}
