package main

import "testing"

const (
	Sunday = iota
	Monday
)

func TestIota(t *testing.T) {
	if Sunday != 0 {
		t.Error("Sunday should be 0")
	}
	if Monday != 1 {
		t.Error("Monday should be 1")
	}
}
