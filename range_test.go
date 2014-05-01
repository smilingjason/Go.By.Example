package main

import (
	"fmt"
	"testing"
)

func TestMapRange(t *testing.T) {
	mymap := make(map[string]string)
	mymap["name"] = "Jason"
	mymap["title"] = "Programmer"

	for key, value := range mymap {
		fmt.Println(key, "=", value)
	}
}
