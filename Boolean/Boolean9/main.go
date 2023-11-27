package main

import "fmt"

func main() {
	var d bool
	var b, a int
	fmt.Scan(&a, &b)
	d = a%2 != 0 || b%2 != 0
	fmt.Println(d)
}
