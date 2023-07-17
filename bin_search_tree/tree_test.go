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

	expected := []any{9, 25, 33, 45, 50, 75, 98}
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

	expected := []any{50, 33, 75, 9, 45, 25, 98}
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

	expected := []any{25, 9, 45, 33, 98, 75, 50}
	fmt.Println("Postorder: ", postOrdered)

	if !reflect.DeepEqual(postOrdered, expected) {
		t.Errorf("expected %v,  got %v", expected, postOrdered)
	}
}

func Test_Delete_LeafNode(t *testing.T) {
	bst := NewTree(intComparator)
	bst.Insert(50)
	fmt.Println("before delete: ", bst.ToArrayInorder())
	bst.Delete(50)
	fmt.Println("after delete: ", bst.ToArrayInorder())
	if bst.Size() != 0 {
		t.Errorf("the size expected is 6, but got %d", bst.Size())
	}
}

func Test_Delete_NodeWithOneChild(t *testing.T) {
	bst := NewTree(intComparator)
	bst.Insert(50)
	bst.Insert(40)
	fmt.Println("before delete: ", bst.ToArrayInorder())
	bst.Delete(50)
	fmt.Println("after delete: ", bst.ToArrayInorder())
	if bst.Size() != 1 {
		t.Errorf("the size expected is 6, but got %d", bst.Size())
	}

	bst = NewTree(intComparator)
	bst.Insert(50)
	bst.Insert(80)
	fmt.Println("before delete: ", bst.ToArrayInorder())
	bst.Delete(50)
	fmt.Println("after delete: ", bst.ToArrayInorder())
	if bst.Size() != 1 {
		t.Errorf("the size expected is 6, but got %d", bst.Size())
	}
}

func Test_Delete_NodeWithBothChildren(t *testing.T) {
	bst := NewTree(intComparator)
	bst.Insert(50)
	bst.Insert(40)
	bst.Insert(80)
	expect := []any{40, 80}
	fmt.Println("before delete: ", bst.ToArrayInorder())
	bst.Delete(50)
	res := bst.ToArrayInorder()
	fmt.Println("after delete: ", res)
	if !reflect.DeepEqual(expect, res) {
		t.Errorf("expected %v, but got %v", expect, res)
	}
	bst = NewTree(intComparator)
	bst.Insert(50)
	bst.Insert(75)
	bst.Insert(33)
	bst.Insert(9)
	bst.Insert(25)
	bst.Insert(98)
	bst.Insert(45)

	expect = []any{9, 25, 33, 45, 75, 98}
	fmt.Println("before delete: ", bst.ToArrayInorder())
	bst.Delete(50)
	res = bst.ToArrayInorder()
	fmt.Println("after delete: ", res)
	if !reflect.DeepEqual(expect, res) {
		t.Errorf("expected %v, but got %v", expect, res)
	}
}

func TestTreeHeight(t *testing.T) {
	bst := NewTree(intComparator)
	bst.Insert(50)
	bst.Insert(75)
	bst.Insert(33)
	bst.Insert(9)
	bst.Insert(25)
	bst.Insert(98)
	bst.Insert(45)

	if bst.Height() != 3 {
		t.Errorf("expected 3, got %d", bst.Height())
	}

	bst = NewTree(intComparator)
	bst.Insert(98)
	bst.Insert(50)
	bst.Insert(45)
	bst.Insert(40)
	bst.Insert(33)
	bst.Insert(25)
	bst.Insert(9)

	if bst.Height() != 6 {
		t.Errorf("expected 6, got %d", bst.Height())
	}
}
