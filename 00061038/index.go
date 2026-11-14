package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var n int
	fmt.Scan(&n)

	l := 1
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		l = l / gcd(l, a) * a
	}
	fmt.Println(l%30 + 1)
}
