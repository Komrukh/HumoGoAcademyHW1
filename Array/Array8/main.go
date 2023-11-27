package main

import (
	"fmt"
	"math"
)

func main() {
	var n, f int
	//fmt.Println("n:")
	fmt.Scan(&n)
	a := make([]int, n)
	//fmt.Println("a[i]:")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	//fmt.Println("f:")
	k := 0
	fmt.Scan(&f)
	for j := 1; j < n; j++ {
		if int(math.Abs(float64(f)-float64(a[k]))) > int(math.Abs(float64(f)-float64(a[j]))) {
			k = j
		}
	}
	fmt.Println(a[k])
}
