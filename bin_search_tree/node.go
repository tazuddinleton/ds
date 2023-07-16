package binsearchtree

// Node is binary search tree node
type Node struct {
	Entry  any
	Parent *Node
	Left   *Node
	Right  *Node
	tree   *Tree
}

func (this *Node) newNode(data any) *Node {
	return &Node{Entry: data, Parent: this, tree: this.tree}
}

func (this *Node) insertLeft(data any) {
	if this.Left == nil {
		this.Left = this.newNode(data)
	} else {
		this.Left.Insert(data)
	}
}

func (this *Node) insertRight(data any) {
	if this.Right == nil {
		this.Right = this.newNode(data)
	} else {
		this.Right.Insert(data)
	}
}

// Insert node to BST
func (this *Node) Insert(data any) {
	if this.tree.compare(data, this.Entry) < 0 {
		this.insertLeft(data)
	} else {
		this.insertRight(data)
	}
}

// Contains checks if a value exists in the tree
func (this *Node) Contains(data any) bool {
	switch cmp := this.tree.compare(data, this.Entry); {
	case cmp == 0:
		return true

	case cmp < 0:
		return this.Left != nil && this.Left.Contains(data)
	default:
		return this.Right != nil && this.Right.Contains(data)
	}
}

// Min returns node containing the minimum value
func (this *Node) Min() any {
	if this.Left == nil {
		return this.Entry
	}
	return this.Left.Min()
}

// Max returns node containing the maximum value
func (this *Node) Max() any {
	if this.Left == nil {
		return this.Entry
	}
	return this.Left.Min()
}

// IsFull returns true if the bst is full (every node contains either 0 or 2 children, not 1)
func (this *Node) IsFull() bool {
	if this.Right != nil && this.Left != nil {
		return this.Left.IsFull() && this.Right.IsFull()
	}

	if this.Right == nil && this.Left == nil {
		return true
	}
	return false
}

func (this *Node) ToArrayInorder() []any {
	items := []any{}
	if this.Left != nil {
		items = append(items, this.Left.ToArrayInorder()...)
	}
	items = append(items, this.Entry)
	if this.Right != nil {
		items = append(items, this.Right.ToArrayInorder()...)
	}
	return items
}

func (this *Node) ToArrayPreorder() []any {
	items := []any{}
	items = append(items, this.Entry)
	if this.Left != nil {
		items = append(items, this.Left.ToArrayPreorder()...)
	}
	if this.Right != nil {
		items = append(items, this.Right.ToArrayPreorder()...)
	}
	return items
}

func (this *Node) ToArrayPostorder() []any {
	items := []any{}
	if this.Left != nil {
		items = append(items, this.Left.ToArrayPostorder()...)
	}
	if this.Right != nil {
		items = append(items, this.Right.ToArrayPostorder()...)
	}
	items = append(items, this.Entry)
	return items
}
