package main

import "fmt"

func main() {
	var a int
	fmt.Println("Enter the number")
	fmt.Scan(&a)
	fmt.Println("Sum: ", (a/10)+(a%10), "\nMul: ", (a/10)*(a%10))
}