package main

import "fmt"

func main() {
	var x1, v1, x2, v2 int64

	fmt.Scan(&x1, &v1, &x2, &v2)

	switch {
	case v1 == v2:
		fmt.Println("WAIT WAIT")

	case (x2-x1)*(v1-v2) > 0:
		fmt.Println("SEE YOU")

	default:
		fmt.Println("BORO BORO")
	}
}
