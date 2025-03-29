package main

import (
    "fmt"
    "math"
    "strconv"
)


func isPrime(num int) bool {
    if num < 2 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}


func generateStrongPasswords(n int, current string, result *[]string) {
    if len(current) == n {
        *result = append(*result, current)
        return
    }
    for i := 1; i <= 9; i++ {
        newNumber := current + strconv.Itoa(i)
        number, _ := strconv.Atoi(newNumber)
        if isPrime(number) {
            generateStrongPasswords(n, newNumber, result)
        }
    }
}

func main() {
    var n int
    fmt.Scan(&n)

    strongPasswords := []string{}
    generateStrongPasswords(n, "", &strongPasswords)

    for _, password := range strongPasswords {
        fmt.Println(password)
    }
}
