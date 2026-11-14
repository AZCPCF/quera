package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	var s string
	if _, err := fmt.Fscan(reader, &n, &s); err != nil || len(s) != 2*n {
		return
	}

	for i := 0; i < n; i++ {
		if s[i] == s[i+n] {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
