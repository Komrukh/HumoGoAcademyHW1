package main

import "fmt"

func main() {
	var a, b, c, k, n int
	fmt.Scan(&a, &b, &c)
	if a > 0 {
		k++
	} else if a < 0 {
		n++
	}
	if b > 0 {
		k++
	} else if b < 0 {
		n++
	}
	if c > 0 {
		k++
	} else if c < 0 {
		n++
	}
	fmt.Println("Количество положительных чисел :", k, "\nКоличество отрецательный чисел :", n)
}
