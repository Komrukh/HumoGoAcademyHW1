package main

import "fmt"

func main() {
	var L int
	fmt.Println("Enter L(cm)")
	fmt.Scan(&L)
	fmt.Println("L(m) : ", L/100)
}
