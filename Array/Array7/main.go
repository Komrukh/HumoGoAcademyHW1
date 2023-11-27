package main

import "fmt"

func main() {
	var n, f int
	//	fmt.Println("n:")
	fmt.Scan(&n)
	a := make([]int, n)
	//	fmt.Println("a[i]:")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	//	fmt.Println("f:")
	k := true
	fmt.Scan(&f)
	for j := 0; j < n; j++ {
		if a[j] == f {
			k = true
		}
	}
	if k == true {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
