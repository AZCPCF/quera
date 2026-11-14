package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 1 {
		fmt.Print(2)
	} else {
		fmt.Print(3)
	}
}
