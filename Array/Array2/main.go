package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	if n >= 1 && n <= 35 {
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}
		for i := 0; i < n/2; i++ {
			a[i], a[n-i-1] = a[n-i-1], a[i]
		}
		fmt.Println(a)
	}
}
