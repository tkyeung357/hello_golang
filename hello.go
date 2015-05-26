package main
import (
    "fmt"
    "time"
    "math/rand"
    )

func main() {
    //print string
    fmt.Println("Hello World")
    
    //time
    fmt.Println("The time is", time.Now())
    
    //math
    fmt.Println("Random Number is:", rand.Intn(10))
    fmt.Println("5 + 10 =", add(5, 10))
    fmt.Println("5.2 + 10 =", add2(5.2, 10))
}
    
func add(x int, y int) int {
    return x+y
}

func add2(x, y int) int {
    return x+y
}