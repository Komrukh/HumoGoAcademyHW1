package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Println("Enter a")
	fmt.Scan(&a)
	fmt.Println("Enter b")
	fmt.Scan(&b)
	fmt.Println("Enter c")
	fmt.Scan(&c)
	fmt.Println("V = ", a*b*c, "\nS = ", 2*(a*b+b*c+a*c))
}
