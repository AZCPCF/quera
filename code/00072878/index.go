package main

import (
    "fmt"
)

func main() {
    var t, a, b int
    fmt.Scan(&t, &a, &b)

    var arrar, mama int
    var time int

    for time < t {
        if time+1 <= t {
            arrar++
            time++
        }
        if time+a <= t {
            time += a
        } else {
            break
        }
        if time+1 <= t {
            mama++
            time++
        }
        if time+b <= t {
            time += b
        } else {
            break
        }
    }

    fmt.Println(arrar, mama)
}
