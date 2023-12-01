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
	l := n - 1
	for k <= l {
		if k == l {
			fmt.Printf("%v ", a[k])
		} else {
			fmt.Printf("%v %v ", a[k], a[l])
		}
		k++
		l--
	}
}

//a[0],a[4],a[1],a[3],a[2]
