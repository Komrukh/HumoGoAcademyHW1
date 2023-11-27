package main

import "fmt"

func main() {
	var a float64
	fmt.Println("Enter a")
	fmt.Scan(&a)
	fmt.Println("V = ", a*a*a, "\nS = ", 6*(a*a))
}
