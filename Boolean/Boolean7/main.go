package main

import "fmt"

func main() {
	var d bool
	var b, a, c int
	fmt.Scan(&a, &b, &c)
	d = a < b && b < c || c < b && b < a
	fmt.Println(d)
}
