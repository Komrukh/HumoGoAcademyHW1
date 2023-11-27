package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	if a < b {
		fmt.Println(b, a)
	} else {
		fmt.Println(a, b)
	}
}
