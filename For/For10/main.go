package main

import "fmt"

func main() {
	var n int
	s := .0
	fmt.Scan(&n)
	for i := 1.; i <= float64(n); i++ {
		s += 1 / i
	}
	fmt.Println(s)
}
