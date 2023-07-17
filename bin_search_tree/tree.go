package binsearchtree

type Tree struct {
	compare func(a, b any) int
	Root    *Node
	size    int
}

func NewTree(cmp func(a, b any) int) *Tree {
	return &Tree{compare: cmp}
}

func (this *Tree) Insert(data any) {
	if this.Root == nil {
		this.Root = &Node{Entry: data, tree: this}
	} else {
		this.Root.Insert(data)
	}
	this.size++
}

func (this *Tree) Delete(data any) {
	if !this.Contains(data) {
		return
	}
	// To delete an item from bst we need to follow the below steps:

	// 1. locate the item
	// 2. if it is a leaf node return nil to the parent
	// 3. if node has one node return it as new node
	// 4. if node has both node find the min node from the right sub-tree then replace
	//      current nodes entry with it, then delete the min from the right sub-tree

	this.Root = deleteNode(this, this.Root, data)
	this.size--
}

func deleteNode(tree *Tree, node *Node, data any) *Node {
	if node.Left == nil && node.Right == nil {
		return nil
	}

	if node.Left == nil {
		return node.Right
	}
	if node.Right == nil {
		return node.Left
	}

	// The item is in the left sub-tree
	if tree.compare(data, node.Entry) < 0 {
		node.Left = deleteNode(tree, node.Left, data)
		return node
	}

	if tree.compare(data, node.Entry) > 0 {
		node.Left = deleteNode(tree, node.Left, data)
		return node
	}

	min := node.Right.Min()
	node.Entry = min
	node.Right = deleteNode(tree, node.Right, min)
	return node
}

func (this *Tree) Contains(data any) bool {
	return this.Root != nil && this.Root.Contains(data)
}

func (this *Tree) ToArrayInorder() []any {
	if this.size == 0 {
		return []any{}
	}
	return this.Root.ToArrayInorder()
}

func (this *Tree) ToArrayPreorder() []any {
	if this.size == 0 {
		return []any{}
	}
	return this.Root.ToArrayPreorder()
}

func (this *Tree) ToArrayPostorder() []any {
	if this.size == 0 {
		return []any{}
	}
	return this.Root.ToArrayPostorder()
}

func (this *Tree) Size() int {
	return this.size
}

func (this *Tree) Height() int {
	return height(this.Root)
}

func height(node *Node) int {
	if node == nil {
		return -1
	}
	if node.Left == nil && node.Right == nil {
		return 0
	}
	return 1 + max(height(node.Left), height(node.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
