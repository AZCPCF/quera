package main

import (
    "fmt"
    "sort"
)

func main() {
    var n int
    fmt.Scan(&n)

    
    numbers := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&numbers[i])
    }

    
    sort.Ints(numbers)

    
    result := make([]int, 0, n)
    left, right := 0, n-1
    for left <= right {
        
        if right >= left {
            result = append(result, numbers[right])
            right--
        }
        
        if left <= right {
            result = append(result, numbers[left])
            left++
        }
    }

    
    for i, val := range result {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(val)
    }
    fmt.Println()
}
