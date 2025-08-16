package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	if t == nil {
		return
	}
	Walk(t.Left,ch)
	ch <- t.Value
	Walk(t.Right,ch)
	
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	c1 := make(chan int,10)
	c2 := make(chan int,10)
	go Walk(t1,c1)
	go Walk(t2,c2)
	var val1 int
	var val2 int
	for i:=0;i<10;i++{
		val1 = <- c1
		val2 = <- c2
		if val1 != val2{
			return false
		}
	}
	return true
}

func main() {
	t1 := tree.New(10)
	t2:= tree.New(12)
	fmt.Print(Same(t1,t2))
}