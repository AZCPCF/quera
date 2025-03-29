package main

import (
    "fmt"
)

func countDreams(a, b, x int) int {
    divisorsA := getDivisors(a)
    divisorsB := getDivisors(b)
    count := 0

    for _, da := range divisorsA {
        for _, db := range divisorsB {
            if da+db <= x {
                count++
            }
        }
    }

    return count
}

func getDivisors(n int) []int {
    var divisors []int
    for i := 1; i <= n; i++ {
        if n%i == 0 {
            divisors = append(divisors, i)
        }
    }
    return divisors
}

func main() {
    var a, b, x int
    fmt.Scan(&a, &b, &x)

    result := countDreams(a, b, x)
    fmt.Println(result)
}
