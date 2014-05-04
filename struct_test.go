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

func TestNewOperation(t *testing.T) {
	//new operator returns a pointer to the T
	var pp *Person = new(Person)
	if pp.name != "" {
		t.Error("Should be empty right after 'new'")
	}
	out(pp)
	if pp.name != "Jason(modified)" {
		t.Error("pp.name should be changed!")
	}
}

func output(p Person) {
	p.name = "Jason(modified)"
}

func out(pp *Person) {
	pp.name = "Jason(modified)"
}
