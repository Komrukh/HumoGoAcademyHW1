package main

import "fmt"

func main() {
	var n int
	j := 0
	fmt.Scan(&n)
	for {
		k := n
		if n != 0 {
			if k < n {
				j++

			}
			fmt.Scan(&n)
		} else {
			break
		}
	}
	fmt.Println(j)
}

/*
func main() {
	var n int
	k := 0
	fmt.Scan(&n)
	for {
		if n != 0 {
			if k < n {
				k = n
			}
			fmt.Scan(&n)
		} else {
			break
		}
	}
	fmt.Println(k)
}
*/
/*
func main() {
	var n int
	j := 0
	fmt.Scan(&n)
	for {
		if n != 0 {
			if n%2 == 0 {
				j++
			}
			fmt.Scan(&n)
		} else {
			break
		}
	}
	fmt.Println(j)
}
*/
/*
func main() {
	var n, s int
	k := .0
	j := 0
	fmt.Scan(&n)
	for {

		if n != 0 {
			s += n
			fmt.Scan(&n)
		} else {
			k = float64(s) / float64(j)
			break
		}
		j++
	}
	fmt.Println(k)
}

*/
/*
func main() {
	var n, s int
	fmt.Scan(&n)
	J:=0
	for {
		if n != 0 {
			s += n
			fmt.Scan(&n)
		} else {
			break
		}
		j++
	}
	fmt.Println(s)
}*/
/*
 3064
func main() {
	var n, k int
	fmt.Scan(&n)
	k = 0
    j:=0
	for {
		if n != 0 {
			fmt.Scan(&n)
			k++
		} else {
			fmt.Scan(&n)
			break
		}
		j++
	}
	fmt.Println(k)
}
*/
