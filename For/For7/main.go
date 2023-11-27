package main

import "fmt"

func main() {
	var n, k, s int
	fmt.Scan(&n, &k)
	for i := n; i <= k; i++ {
		s += i
	}
	fmt.Println(s)
}
