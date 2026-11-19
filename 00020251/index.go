package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)

	var n int
	fmt.Fscan(reader, &n)

	g := make([][]int, n+1)
	for i := 2; i <= n; i++ {
		var p int
		fmt.Fscan(reader, &p)
		g[p] = append(g[p], i)
	}

	depth := make([]int64, n+1)
	depth[1] = 1
	queue := make([]int, 0, n)
	queue = append(queue, 1)

	var ans int64
	for head := 0; head < len(queue); head++ {
		u := queue[head]
		ans += depth[u]
		for _, v := range g[u] {
			depth[v] = depth[u] + 1
			queue = append(queue, v)
		}
	}

	fmt.Println(ans)
}
