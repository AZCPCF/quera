package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            break
        }
        n, _ := strconv.Atoi(strings.TrimSpace(line))
        square := n * n
        reversed := reverseString(strconv.Itoa(square))
        fmt.Fprintln(writer, reversed)
    }
}
