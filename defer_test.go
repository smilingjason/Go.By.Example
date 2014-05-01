package main

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	var pi *int // pointer to int
	var a int
	pi = &a
	fmt.Println("*pi = ", *pi)
	defer deferredOutput(pi, t) // defer to execute before turn
	a = 2
	fmt.Println("*pi = ", *pi)
}

func deferredOutput(pi *int, t *testing.T) {
	if *pi != 2 {
		t.Errorf("expetected 0, but %d", *pi)
	}
}
