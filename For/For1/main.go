package main

import "fmt"

func main() {
	var n int
	var k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	for i := 1; i <= n; i++ {
		fmt.Println(k)
	}
}
