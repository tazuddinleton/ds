package binsearchtree

import "fmt"

type Tree struct {
	compare func(a, b any) int
	Root    *Node
}

func NewTree(cmp func(a, b any) int) *Tree {
	return &Tree{compare: cmp}
}

func (this *Tree) Insert(data any) {
	if this.Root == nil {
		fmt.Println("inserting root node")
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
