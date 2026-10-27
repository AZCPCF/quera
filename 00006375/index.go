package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b, c float64
	fmt.Fscan(reader, &a)
	fmt.Fscan(reader, &b)
	fmt.Fscan(reader, &c)
	sum := a + b + c

	switch {
	case a == b && b == c:
		fmt.Print(0)
	case 3*a == sum ||
		3*b == sum ||
		3*c == sum:
		fmt.Print(1)
	default:
		fmt.Print(2)
	}
}
