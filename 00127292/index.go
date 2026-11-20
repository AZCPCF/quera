package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	writer := bufio.NewWriterSize(os.Stdout, 1<<20)
	defer writer.Flush()

	readInt := func() int {
		n := 0
		c, _ := reader.ReadByte()
		for c == ' ' || c == '\n' || c == '\r' {
			c, _ = reader.ReadByte()
		}
		neg := false
		if c == '-' {
			neg = true
			c, _ = reader.ReadByte()
		}
		for c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
			c, _ = reader.ReadByte()
		}
		if neg {
			return -n
		}
		return n
	}

	t := readInt()
	for ; t > 0; t-- {
		m := readInt()
		b := make([]int, m)
		for i := range b {
			b[i] = readInt()
		}
		pi := make([]int, m)
		for i := 1; i < m; i++ {
			j := pi[i-1]
			for j > 0 && b[i] != b[j] {
				j = pi[j-1]
			}
			if b[i] == b[j] {
				j++
			}
			pi[i] = j
		}
		p := m - pi[m-1]
		writer.WriteString(strconv.Itoa(p))
		writer.WriteByte('\n')
		for i := 0; i < p; i++ {
			if i > 0 {
				writer.WriteByte(' ')
			}
			writer.WriteString(strconv.Itoa(b[i]))
		}
		writer.WriteByte('\n')
	}
}
