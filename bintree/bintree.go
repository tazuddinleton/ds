package bintree

type Node[T any] struct {
	Data  T
	Left  *Node[T]
	Right *Node[T]
}

func TreeNode[T any](data T, left, right *Node[T]) *Node[T] {
	return &Node[T]{Data: data, Left: left, Right: right}
}

// Inorder traversal (left, root, right)
func InOrdered[T any](tree *Node[T], fn func(T)) {
	if tree == nil {
		return
	}

	if tree.Left != nil {
		InOrdered(tree.Left, fn)
	}
	fn(tree.Data)
	if tree.Right != nil {
		InOrdered(tree.Right, fn)
	}
}

// PreOrdered traversal (root, left, right)
func PreOrdered[T any](tree *Node[T], fn func(T)) {
	if tree == nil {
		return
	}
	fn(tree.Data)
	if tree.Left != nil {
		PreOrdered(tree.Left, fn)
	}
	if tree.Right != nil {
		PreOrdered(tree.Right, fn)
	}
}

// PostOrdred traversal (left, right, root)
func PostOrdered[T any](tree *Node[T], fn func(T)) {
	if tree == nil {
		return
	}
	if tree.Left != nil {
		PostOrdered(tree.Left, fn)
	}
	if tree.Right != nil {
		PostOrdered(tree.Right, fn)
	}
	fn(tree.Data)
}
