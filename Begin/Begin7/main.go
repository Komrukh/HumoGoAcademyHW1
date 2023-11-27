package main

import "fmt"

func main() {
	var R float64
	pi := 3.14
	fmt.Println("Enter R")
	fmt.Scan(&R)
	fmt.Println("L = ", 2*pi*R, "\nS = ", pi*(R*R))
}
