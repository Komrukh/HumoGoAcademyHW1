package main

import "fmt"

func main() {
	var n, k int
	s := 0
	fmt.Scan(&n, &k)
	for i := n; i <= k; i++ {
		s += i * i
	}
	fmt.Println(s)
}
