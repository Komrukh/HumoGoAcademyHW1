package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Println("Enter a")
	fmt.Scan(&a)
	fmt.Println("Enter b")
	fmt.Scan(&b)
	fmt.Println("Middle geometric : ", math.Sqrt(a*b))
}
