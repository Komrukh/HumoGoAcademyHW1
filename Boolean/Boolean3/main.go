package main

import "fmt"

func main() {
	var a bool
	var b int
	fmt.Scan(&b)
	a = b%2 == 0
	fmt.Println(a)
}
