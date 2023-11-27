package main

import "fmt"

func main() {
	var c bool
	var b, a int
	fmt.Scan(&a, &b)
	c = (a > 2) && (b <= 3)
	fmt.Println(c)
}
