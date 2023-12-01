package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	k := a[0]
	l := -a[0]
	for j := 1; j < n; j++ {
		if a[j] > 0 && a[j] < k {
			k = a[j]
		}
		if a[j] < 0 && a[j] > l {
			l = a[j]
		}
	}
	for i := 0; i < n-2; i++ {
		if a[i] == k {
			a = append(a[:i], a[i+1:]...)
		}
		if a[i] == l {
			a = append(a[:i], a[i+1:]...)
		}
	}
	fmt.Println(a)
}
