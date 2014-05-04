package main

import (
	"fmt"
	"testing"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func TestTreeTranversalInorder(t *testing.T) {
	fmt.Println("Hello")
	//prepare tree
	root := initTree2()
	//prepare channel to communicate
	c := make(chan int, 20)
	//tranverse tree
	go func() {
		inorder(root, c)
		close(c)
	}()
	//gather tranverse result
	for v := range c {
		fmt.Println(v)
	}
}

func initTree() *Tree {
	root := &Tree{Left: nil, Value: 3, Right: nil}
	root.Left = &Tree{Left: nil, Value: 1, Right: nil}
	root.Right = &Tree{Left: nil, Value: 8, Right: nil}
	root.Left.Left = &Tree{Left: nil, Value: 1, Right: nil}
	root.Left.Right = &Tree{Left: nil, Value: 2, Right: nil}
	root.Right.Left = &Tree{Left: nil, Value: 5, Right: nil}
	root.Right.Right = &Tree{Left: nil, Value: 13, Right: nil}
	/*
	       3
	      /  \
	     1    8
	    / \  / \
	   1   2 5 13
	*/
	return root
}

func initTree2() *Tree {
	root := &Tree{Left: nil, Value: 3, Right: nil}
	insert(root, &Tree{Left: nil, Value: 1, Right: nil})
	insert(root, &Tree{Left: nil, Value: 8, Right: nil})
	insert(root, &Tree{Left: nil, Value: 1, Right: nil})
	insert(root, &Tree{Left: nil, Value: 2, Right: nil})
	insert(root, &Tree{Left: nil, Value: 13, Right: nil})
	insert(root, &Tree{Left: nil, Value: 5, Right: nil})

	return root
}

func insert(root *Tree, child *Tree) {
	if child.Value < root.Value {
		if root.Left == nil {
			root.Left = child
		} else {
			insert(root.Left, child)
		}
	} else {
		if root.Right == nil {
			root.Right = child
		} else {
			insert(root.Right, child)
		}
	}
}

func inorder(root *Tree, c chan int) {
	if root == nil {
		return
	}
	inorder(root.Left, c)
	// visit this node
	c <- root.Value
	inorder(root.Right, c)
}
