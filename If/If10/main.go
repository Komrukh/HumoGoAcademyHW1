package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a != b {
		fmt.Println("A =", a+b, "B =", b+a)
	} else {
		fmt.Println("A =", a-a, "B =", b-b)
	}
}
