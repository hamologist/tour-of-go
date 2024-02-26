package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var inner func(t *tree.Tree)
	inner = func(t *tree.Tree) {
		if t == nil {
			return
		}

		inner(t.Left)
		ch <- t.Value
		inner(t.Right)
	}
	inner(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1Ch, t2Ch := make(chan int), make(chan int)

	go Walk(t1, t1Ch)
	go Walk(t2, t2Ch)
	for {
		v1, ok1 := <-t1Ch
		v2, ok2 := <-t2Ch

		if v1 != v2 || ok1 != ok2 {
			return false
		}

		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

