package main

import "fmt"

func main() {
	var a, b float64
	fmt.Println("Enter a")
	fmt.Scan(&a)
	fmt.Println("Enter b")
	fmt.Scan(&b)
	fmt.Println("Middle arithmetic : ", (a+b)/2)
}
