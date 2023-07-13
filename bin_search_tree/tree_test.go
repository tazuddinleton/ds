package binsearchtree

import (
	"fmt"
	"testing"
)

func intComparator(a, b any) int {
	assertA := a.(int)
	assertB := b.(int)
	if assertA < assertB {
		return -1
	}
	if assertA == assertB {
		return 0
	}
	return 1
}

func TestTree(t *testing.T) {
	bst := NewTree(intComparator)
	bst.Insert(50)
	fmt.Println("bst node inserted")
	if bst.Root == nil || bst.Root.Entry != 50 {
		t.Errorf("wanted Root node to be 50, but got %v", bst.Root)
	}
}
