package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	if a < b {
		fmt.Println("A =", a, "B =", b)
	} else {
		fmt.Println("A =", b, "B =", a)
	}
}
