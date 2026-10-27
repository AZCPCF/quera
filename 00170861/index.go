package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var T int
	fmt.Fscan(reader, &T)

	for ; T > 0; T-- {
		var N int
		fmt.Fscan(reader, &N)

		height := make([]int, N)

		for i := 0; i < N; i++ {
			fmt.Fscan(reader, &height[i])
		}

		answer := 0

		for i := 0; i < N; i++ {
			var M int
			fmt.Fscan(reader, &M)

			for j := 0; j < M; j++ {
				var neighbor int
				fmt.Fscan(reader, &neighbor)

				neighbor--

				diff := abs(height[i] - height[neighbor])

				if diff > answer {
					answer = diff
				}
			}
		}

		fmt.Fprintln(writer, answer)
	}
}
