package main

import "fmt"

func main() {
	var a int
	fmt.Println("Enter the number")
	fmt.Scan(&a)
	fmt.Println("Hundreds: ", a/100)
}
