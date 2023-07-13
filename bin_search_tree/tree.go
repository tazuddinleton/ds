package binsearchtree

type Comparator interface {
	Compare(a, b any) int
}

type Tree struct {
	comparator Comparator
	Root       *Node
}

func (this *Tree) compare(a, b any) int {
	return this.comparator.Compare(a, b)
}

func NewTree(cmp Comparator) *Tree {
	return &Tree{comparator: cmp}
}

func (this *Tree) Add(data any) {
	if this.Root == nil {
		this.Root = &Node{Entry: data, tree: this}
	} else {
		this.Root.Add(data)
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
