package main

import "fmt"


var daysInMonth = []int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}

func calculateDays(m1, d1, m2, d2 int) int {
    
    totalDays1, totalDays2 := d1, d2

    for i := 0; i < m1-1; i++ {
        totalDays1 += daysInMonth[i]
    }
    for i := 0; i < m2-1; i++ {
        totalDays2 += daysInMonth[i]
    }

    
    if totalDays1 > totalDays2 {
        return totalDays1 - totalDays2
    }
    return totalDays2 - totalDays1
}

func main() {
    var m1, d1, m2, d2 int
    fmt.Scan(&m1, &d1, &m2, &d2)

    result := calculateDays(m1, d1, m2, d2)
    fmt.Println(result)
}
