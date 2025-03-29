package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    
    var n int
    fmt.Scan(&n)

    
    names := make([]string, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(reader, &names[i])
    }

    
    for i := 1; i < n; i++ {
        for j := i - 1; j >= 0; j-- {
            fmt.Printf("%s: salam %s!\n", names[i], names[j])
        }
    }

    
    for i := 0; i < n; i++ {
        fmt.Printf("%s: khodafez bacheha!\n", names[i])
        for j := i + 1; j < n; j++ {
            fmt.Printf("%s: khodafez %s!\n", names[j], names[i])
        }
    }
}
