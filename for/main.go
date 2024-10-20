
package main

import "fmt"

func main() {
    i := 1

    // Correct way of incrementing 'i'
    for i <= 6 {
        fmt.Println(i)
        i++ // Standalone increment in Go
    }

    // Traditional for-loop works as expected
    for j := 0; j < 9; j++ {
        fmt.Println(j)
    }

    // If you want to iterate over a range of numbers, use a slice or array
    for k := 0; k < 6; k++ {
        fmt.Println("range: ", k)
    }

    // Infinite loop with a break
    for {
        fmt.Println("infinite")
        break
    }

    // Using range to iterate over an array
    numbers := []int{0, 1, 2, 3, 4, 5, 6} // Array or slice needed
    for n := range numbers {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
