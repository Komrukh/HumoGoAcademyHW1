package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if float64(arr[i]) == math.Sqrt(float64(arr[j])) {
				arr = append(arr[:i], arr[i+1:]...)
				n--
			}
		}
	}
	fmt.Println(arr)
}
