package main

import "fmt"

func main() {
	var n, k int
	s := 1
	fmt.Scan(&n, &k)
	for i := n; i <= k; i++ {
		s *= i
	}
	fmt.Println(s)
}
