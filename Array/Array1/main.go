package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var t bool
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 1; i < n; i++ {
		if a[i] > 0 && a[i-1] > 0 || a[i] < 0 && a[i-1] < 0 {
			t = true
			break
		} else {
			t = false
		}
	}
	if t == true {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
