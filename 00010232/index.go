package main

import (
    "fmt"
)

func main() {
    var n, l int
    fmt.Scan(&n, &l)

    redLights := make([][3]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&redLights[i][0], &redLights[i][1], &redLights[i][2]) 
    }

    time := 0
    currentPosition := 0

    for _, light := range redLights {
        d, r, g := light[0], light[1], light[2]

        
        time += d - currentPosition
        currentPosition = d

        
        cycle := time % (r + g)
        if cycle < r {
            
            time += r - cycle
        }
    }

    
    time += l - currentPosition

    fmt.Println(time)
}
