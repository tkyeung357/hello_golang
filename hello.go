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
    fmt.Println("51 + 10 =", add2(51, 10))
    
    //multi results, tour 7
    a, b := swap("hello", "world");
    fmt.Println(a, b)
    
    //named return values
    fmt.Println(split(17))
}
    
func add(x int, y int) int {
    return x+y
}

func add2(x, y int) int {
    return x+y
}

func swap(x, y string) (string, string) {
    return x, y
}

func split(sum int) (x, y int) {
    
    x = sum * 4 / 9
    y = sum - x
    return
}