package main

import "fmt"

func main() {
	var a, b, c, k int
	fmt.Scan(&a, &b, &c)
	if a > 0 {
		k++
	}
	if b > 0 {
		k++
	}
	if c > 0 {
		k++
	}
	fmt.Println("Количество положительных чисел :", k)
}
