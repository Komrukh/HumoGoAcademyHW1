package main

import (
	"fmt"
)

func main() {
	var n, a, s int
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		s += (n + a) * (n + a)
		a = a + 1
	}
	fmt.Println(s)
}
