package main

import "fmt"

func fibonacci(n int) {
    a, b := 0, 1

    for i := 0; i < n; i++ {
        fmt.Print(a, " ")

        a, b = b, a+b
    }
}

func main() {
    fmt.Print("Enter the number of terms: ")
    var n int
    fmt.Scan(&n)

    fmt.Println("Fibonacci Series:")
    fibonacci(n)
}
