package main

import "fmt"

func main() {
	var n, i float32
	fmt.Scan(&n)
	for i = 1.2; i <= 2.1; i += .2 {
		fmt.Printf("%v\n", n*i)
	}
}
