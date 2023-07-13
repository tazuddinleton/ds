package binsearchtree

import "testing"

func TestNew(t *testing.T) {
	root := New[int64](10)

	if root == nil {
		t.Errorf("wanted a valid tree node, but got nil")
	}
}
