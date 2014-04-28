package main

import "fmt"
import "github.com/smilingjason/mylib"
import "github.com/smilingjason/mylib/math"
import "github.com/smilingjason/util"
import utilmath "github.com/smilingjason/util/math"

func main() {
    fmt.Printf("Hello, world. 世界\n")
    mylib.Hello()
    fmt.Println(add(1, 1))
    fmt.Println(math.Add(10, 10))
    fmt.Println(math.Sub(10, 10))
    fmt.Printf("Hello %s\n", new(World))
    fmt.Println(new(World))
    fmt.Println(util.Hi())
    fmt.Println(util.Hi2())
    fmt.Println(utilmath.Add(2, 2))
}

func add(x int, y int) int {
    for i := 0; i < 100; i++ {
    }
    return x + y
}

type World struct{}

func (w *World) String() string {
    return "世界"
}
