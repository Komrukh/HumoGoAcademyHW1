package main

import "fmt"

func main() {
	var a int
	fmt.Println("Enter the number")
	fmt.Scan(&a)
	fmt.Println("Tens: ", a/10, "\nUnits: ", a%10)
}
