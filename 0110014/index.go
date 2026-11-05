package main

import "fmt"

func main() {
    var left1, right1, left2, right2 string

    fmt.Scan(&left1, &right1)
    fmt.Scan(&left2, &right2)

    if left1 == right1 || left1 == right2 ||
        left2 == right1 || left2 == right2 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
