package bintree

import (
	"reflect"
	"testing"
)

func TestMakeTree(t *testing.T) {
	tree := TreeNode(1, TreeNode(2, nil, nil), TreeNode(3, nil, nil))

	if tree.Data != 1 {
		t.Errorf("wanted 1, got %d", tree.Data)
	}
	if tree.Left.Data != 2 {
		t.Errorf("wanted 2, got %d", tree.Left.Data)
	}
	if tree.Right.Data != 3 {
		t.Errorf("wanted 3, got %d", tree.Right.Data)
	}
}

func TestInorderedTraversal(t *testing.T) {
	tree := TreeNode(1, TreeNode(2, TreeNode(4, nil, nil), TreeNode(5, nil, nil)), TreeNode(3, nil, nil))
	// 			1
	// 		2		3
	// 	4		5

	arr := []int{}
	InOrdered(tree, func(item int) { arr = append(arr, item) })

	if !reflect.DeepEqual(arr, []int{4, 2, 5, 1, 3}) {
		t.Errorf("wanted [4, 2, 5, 1, 3], got %v", arr)
	}
}

func TestPreOrderdTraversal(t *testing.T) {
	tree := TreeNode(1, TreeNode(2, TreeNode(4, nil, nil), TreeNode(5, nil, nil)), TreeNode(3, nil, nil))
	// 			1
	// 		2		3
	// 	4		5

	arr := []int{}
	PreOrdered(tree, func(item int) { arr = append(arr, item) })

	if !reflect.DeepEqual(arr, []int{1, 2, 4, 5, 3}) {
		t.Errorf("wanted [1, 2, 3, 4, 5], got %v", arr)
	}
}

func TestPostOrderdTraversal(t *testing.T) {
	tree := TreeNode(1, TreeNode(2, TreeNode(4, nil, nil), TreeNode(5, nil, nil)), TreeNode(3, nil, nil))
	// 			1
	// 		2		3
	// 	4		5

	arr := []int{}
	PostOrdered(tree, func(item int) { arr = append(arr, item) })

	if !reflect.DeepEqual(arr, []int{4, 5, 2, 3, 1}) {
		t.Errorf("wanted [4, 5, 2, 3, 1], got %v", arr)
	}
}
