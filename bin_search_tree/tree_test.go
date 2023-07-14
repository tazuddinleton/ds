package binsearchtree

import (
	"fmt"
	"reflect"
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

/*
                  50

                      75
          33
   9             45          98

      25


*/

func TestNewTree(t *testing.T) {
	bst := NewTree(intComparator)
	bst.Insert(50)
	fmt.Println("bst node inserted")
	if bst.Root == nil || bst.Root.Entry != 50 {
		t.Errorf("wanted Root node to be 50, but got %v", bst.Root)
	}
}

func TestToArrayInorder(t *testing.T) {
	bst := NewTree(intComparator)
	bst.Insert(50)
	bst.Insert(75)
	bst.Insert(33)
	bst.Insert(9)
	bst.Insert(25)
	bst.Insert(98)
	bst.Insert(45)

	inOrdered := bst.ToArrayInorder()
	if len(inOrdered) != 7 {
		t.Errorf("wanted 7 items but got %d", len(inOrdered))
	}

	expected := []int{9, 25, 33, 45, 50, 75, 98}
	fmt.Println("Inorder: ", inOrdered)

	if !reflect.DeepEqual(inOrdered, expected) {
		t.Errorf("expected %v,  got %v", expected, inOrdered)
	}
}

func TestToArrayPreorder(t *testing.T) {
	bst := NewTree(intComparator)
	bst.Insert(50)
	bst.Insert(75)
	bst.Insert(33)
	bst.Insert(9)
	bst.Insert(25)
	bst.Insert(98)
	bst.Insert(45)

	preOrdered := bst.ToArrayPreorder()
	if len(preOrdered) != 7 {
		t.Errorf("wanted 7 items but got %d", len(preOrdered))
	}

	expected := []int{50, 33, 75, 9, 45, 25, 98}
	fmt.Println("Preorder: ", preOrdered)

	if !reflect.DeepEqual(preOrdered, expected) {
		t.Errorf("expected %v,  got %v", expected, preOrdered)
	}
}

func TestToArrayPostorder(t *testing.T) {
	bst := NewTree(intComparator)
	bst.Insert(50)
	bst.Insert(75)
	bst.Insert(33)
	bst.Insert(9)
	bst.Insert(25)
	bst.Insert(98)
	bst.Insert(45)

	postOrdered := bst.ToArrayPostorder()
	if len(postOrdered) != 7 {
		t.Errorf("wanted 7 items but got %d", len(postOrdered))
	}

	expected := []int{25, 9, 45, 33, 98, 75, 50}
	fmt.Println("Postorder: ", postOrdered)

	if !reflect.DeepEqual(postOrdered, expected) {
		t.Errorf("expected %v,  got %v", expected, postOrdered)
	}
}
