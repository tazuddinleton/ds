package binsearchtree

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
		this.Left.Add(data)
	}
}

func (this *Node) insertRight(data any) {
	if this.Right == nil {
		this.Right = this.newNode(data)
	} else {
		this.Right.Add(data)
	}
}

func (this *Node) Add(data any) {
	if this.tree.comparator.Compare(data, this.Entry) < 0 {
		this.insertLeft(data)
	} else {
		this.insertRight(data)
	}
}

func (this *Node) Delete(data any) {
	if !this.Contains(data) {
		return
	}
}

func (this *Node) Contains(data any) bool {
	switch cmp := this.tree.compare(this.Entry, data); {
	case cmp == 0:
		return true

	case cmp < 0:
		return this.Left != nil && this.Left.Contains(data)
	default:
		return this.Right != nil && this.Right.Contains(data)
	}
}

func (this *Node) Min() any {
	if this.Left == nil {
		return this.Entry
	}
	return this.Left.Min()
}

func (this *Node) Max() any {
	if this.Left == nil {
		return this.Entry
	}
	return this.Left.Min()
}
