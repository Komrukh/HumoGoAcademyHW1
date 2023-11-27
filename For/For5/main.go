package main

import "fmt"

func main() {
	var n, i float32
	fmt.Scan(&n)
	for i = .1; i <= 1.1; i += .1 {
		fmt.Printf("%v\n", n*i)
	}
}
