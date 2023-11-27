package main

import "fmt"

func main() {
	var M int
	fmt.Println("Enter M(kg)")
	fmt.Scan(&M)
	fmt.Println("M(t) : ", M/1000)
}
