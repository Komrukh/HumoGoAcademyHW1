package main

import "fmt"

func main() {
	var a, b float64
	fmt.Println("Enter a")
	fmt.Scan(&a)
	fmt.Println("Enter b")
	fmt.Scan(&b)
	fmt.Println("Sum : ", (a*a)+(b*b), "\nDif:", (a*a)-(b*b), "\nMul:", (a*a)*(b*b), "\nDiv:", (a*a)/(b*b))
}
