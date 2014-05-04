package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	fmt.Println("Hello")
	var mychan = make(chan int, 10)
	go send(mychan)
	go receive(mychan)
	time.Sleep(30 * time.Second)
}

func send(c chan int) {
	fmt.Println("Sleeping in send v2")
	time.Sleep(3 * time.Second)
	c <- 100
	fmt.Println("Data sent to chan")
}
func receive(c chan int) {
	fmt.Println("going to receive data...")
	time.Sleep(30 * time.Second)
	result := <-c
	fmt.Println("received:", result)
}
