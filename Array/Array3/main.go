package main

import "fmt"

func main() {
	var n int
	var s int
	fmt.Scan(&n)
	a := make([]int, n)
	if n >= 1 && n <= 35 {
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}
		for i := 1; i < n; i++ {
			if a[i] > a[i-1] {
				s = a[i]
			}
		}
		fmt.Println(s)
	}
}
