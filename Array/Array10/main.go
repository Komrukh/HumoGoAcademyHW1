package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	mi := a[0]
	ma := a[0]
	for j := 1; j < n; j++ {
		if a[j] < mi {
			mi = a[j]
		}
		if a[j] > ma {
			ma = a[j]
		}
	}
	for i := 0; i < n; i++ {
		if a[i] == ma {
			a[i] = mi
		}
		fmt.Printf("%v ", a[i])
	}
}
