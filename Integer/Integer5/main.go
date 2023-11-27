package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Enter a")
	fmt.Scan(&a)
	fmt.Println("Enter b")
	fmt.Scan(&b)
	fmt.Println("Length of the unoccupied part: ", a-(a/b)*b)
}
