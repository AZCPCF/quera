package main

import "fmt"

func main() {
	var canada, france, years int
	fmt.Scan(&canada, &france, &years)
	if (canada+years)*2 >= france {
		fmt.Print("YES")

	} else {
		fmt.Print("NO")
	}
}
