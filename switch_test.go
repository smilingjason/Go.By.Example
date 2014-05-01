package main

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	var a int = 1
	switch a {
	case 1:
		fmt.Println("OK")
	default:
		t.Error("Should not come to this branch")
	}
}
