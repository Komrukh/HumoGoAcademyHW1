package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	m = a[0]
	mi = i
	for i := 0; i < n; i++ {
		if m == 0 {
		}
	}
	fmt.Println(a)

}

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	a := make([]int, n)
// 	k := 0
// 	//var a []int
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&a[i])
// 	}
// 	for j := 1; j < n; j++ {
// 		if a[j] > a[j-1] {
// 			k++
// 		}
// 	}
// 	fmt.Printf("%v ", k)
// }

/*
func main() {
	var n int
	var a int
	var res bool
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a)
		if a == 0 {
			res = true
		}
	}
	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
*/
/*
func main() {
	var n int
	var a int
	k := 0
	l := 0
	m := 0
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a)
		if a == 0 {
			k++
		} else if a > 0 {
			l++
		} else {
			m++
		}
	}
	fmt.Println("Количество 0: ", k, l, m)
}*/

/*func main() {
	var a,b,c,d int
	fmt.Scan(&a,&b,&c,&d)
	for i := a; i < b; i++ {
		if i%d==c{
			fmt.Println(i)
		}
	}
}

func main() {
	var w string
	var d int
	fmt.Scan(&w)
	d = len(w)
	if d > 10 {
		fmt.Printf("%v%v%v", string(w[0]), d-2, string(w[d-1]))
	} else {
		fmt.Println(w)
	}
}
func main() {
	var s string
	fmt.Scan(&s)
	var n int
	n = len(s)
	fmt.Printf("%v\n%v\n%v\n%v", string(s[2]), string(s[n-1]), string(s[:5]), string(s[:n-2]))
}*/
