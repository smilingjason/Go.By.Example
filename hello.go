package main

import "fmt"

func main() {
	fmt.Printf("Hello, world. 世界\n")
	fmt.Println(add(1, 1))
	fmt.Printf("Hello, world. %s\n", new(World))
}

func add(x int, y int) int {
	return x + y
}

type World struct{}

func (w *World) String() string {
	return "世界"
}
