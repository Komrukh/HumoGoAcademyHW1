package main

import "fmt"

func main() {
	var n, k, f int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	f = n / k
	if n%k == 0 {
		f++
	}
	for i := 0; i < f; i++ {
		for j := i * k; j < (i+1)*k && j < n; j++ {
			fmt.Print(a[j], " ")
		}
	}
}
