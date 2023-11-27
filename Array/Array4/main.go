package main

import "fmt"

func main() {
	var n int
	var s int
	fmt.Scan(&n)
	a := make([]int, n)
	if n >= 1 && n <= 100 {
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}
		for i := 1; i < n-1; i++ {
			if a[i] > a[i-1] && a[i] > a[i+1] {
				s++
			}
		}
		fmt.Println(s)
	}
}
