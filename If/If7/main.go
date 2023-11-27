package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	if a < b {
		fmt.Println("Порядковый номер меньшего :", 1)
	} else {
		fmt.Println("Порядковый номер меньшего :", 2)
	}
}
