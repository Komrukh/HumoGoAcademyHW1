package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a > b {
		fmt.Println("Max of number :", a)
	} else {
		fmt.Println("Max of number :", b)
	}
}
