package main

import "fmt"

func main() {
	var n, k float64
	fmt.Scan(&n)
	for i := 1; i <= 10; i++ {
		k = n * float64(i)
		fmt.Println(k)
	}
}
