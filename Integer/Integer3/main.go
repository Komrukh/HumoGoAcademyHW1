package main

import "fmt"

func main() {
	var B int
	fmt.Println("Enter B(byte)")
	fmt.Scan(&B)
	fmt.Println("B(kbyte) : ", B/1024)
}
