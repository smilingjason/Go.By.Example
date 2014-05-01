package main

import "testing"

func TestStructLiteral(t *testing.T) {
	var me = Person{"Jason", 40}
	if me.name != "Jason" {
		t.Error("name not set correctly")
	}
	me.name = "Jason Huang"
	if me.name != "Jason Huang" {
		t.Error("name not changed correctly")
	}
	output(me)
	if me.name != "Jason Huang" {
		t.Error("name should not be changed")
	}
	out(&me)
	if me.name != "Jason(modified)" {
		t.Error("name should not be changed")
	}
}

func output(p Person) {
	p.name = "Jason(modified)"
}

func out(pp *Person) {
	pp.name = "Jason(modified)"
}
