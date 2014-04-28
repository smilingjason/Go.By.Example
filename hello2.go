package main

import "fmt"

//import "github.com/smilingjason/mylib"
//import "github.com/smilingjason/mylib/math"
import "github.com/smilingjason/util"
import utilmath "github.com/smilingjason/util/math"

func mymain() {
    fmt.Printf("Hello2 %s\n", new(World))
    fmt.Println(new(World))
    fmt.Println(util.Hi())
    fmt.Println(util.Hi2())
    fmt.Println(utilmath.Add(2, 2))

}
