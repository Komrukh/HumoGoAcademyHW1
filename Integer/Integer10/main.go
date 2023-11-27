package main

import "fmt"

func main() {
	var a int
	fmt.Println("Enter the number")
	fmt.Scan(&a)
	fmt.Println("Ones: ", (a % 10), "\nTens: ", (a/10)%10)
}
