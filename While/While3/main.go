package main

import "fmt"

func main() {
	var k, m, u, r int
	fmt.Scan(&k, &m)
	if k > m {
		for k > m {
			k -= m
			r = k
			u++
		}
		fmt.Println(u)
		fmt.Println(r)
	} else {
		r = k
		//u = 0
		fmt.Println(u)
		fmt.Println(r)
	}

}
