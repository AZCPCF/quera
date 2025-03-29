package main

import (
    "fmt"
    "math"
)


func Factorial(n int) float64 {
    if n == 0 {
        return 1
    }
    result := 1.0
    for i := 1; i <= n; i++ {
        result *= float64(i)
    }
    return result
}


func Exponential(x float64, n int) float64 {
    result := 0.0
    for i := 0; i < n; i++ {
        result += math.Pow(x, float64(i)) / Factorial(i)
    }
    return result
}

func main() {
    var x float64
    var n int

    
    fmt.Scan(&x)
    fmt.Scan(&n)

    
    result := Exponential(x, n)

    
    fmt.Printf("%.3f\n", result)
}
