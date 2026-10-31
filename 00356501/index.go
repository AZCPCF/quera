package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var a, b, c, d, e, f int
		fmt.Scan(&a, &b, &c, &d, &e, &f)
		if a+b+c == d+e+f {
			fmt.Println("Possible")
		} else {
			fmt.Println("Impossible")
		}
	}
}
