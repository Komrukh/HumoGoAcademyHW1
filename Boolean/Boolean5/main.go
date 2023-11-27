package main

import "fmt"

func main() {
	var c bool
	var b, a int
	fmt.Scan(&a, &b)
	c = a >= 0 && b < -2
	fmt.Println(c)
}
