package binsearchtree

type Tree struct {
	compare func(a, b any) int
	Root    *Node
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
}

func (this *Tree) Delete(data any) {
	if !this.Contains(data) {
		return
	}
	this.Root.Delete(data)
}

func (this *Tree) Contains(data any) bool {
	return this.Root != nil && this.Root.Contains(data)
}

func (this *Tree) ToArrayInorder() []any {
	return this.Root.ToArrayInorder()
}

func (this *Tree) ToArrayPreorder() []any {
	return this.Root.ToArrayPreorder()
}

func (this *Tree) ToArrayPostorder() []any {
	return this.Root.ToArrayPostorder()
}
