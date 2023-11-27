package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	k := 0
	for i := 0; i < n; i++ {
		if a[i]%2 != 0 {
			k = a[i]
			break
		}
	}
	for i := 0; i < n; i++ {
		if a[i]%2 != 0 {
			if a[i] < k {
				k = a[i]
			}
		}
	}

	fmt.Println(k)
}
