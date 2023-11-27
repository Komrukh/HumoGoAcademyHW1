package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Enter a")
	fmt.Scan(&a)
	fmt.Println("Enter b")
	fmt.Scan(&b)
	fmt.Println("Count of segments : ", a/b)
}
