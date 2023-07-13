package objectpointer

import (
	"fmt"
	"testing"
)

func TestResetFree(t *testing.T) {
	resetFree()

	fmt.Println(NEXT)
	if FREE == -1 {
		t.Errorf("wanted 1, got -1")
	}
}

func TestAllocateObject(t *testing.T) {
	resetFree()

	allocateObject(9)
	allocateObject(25)
	allocateObject(13)
	allocateObject(13)
	allocateObject(13)
	allocateObject(13)
	allocateObject(13)
	allocateObject(13)
	allocateObject(10)
	allocateObject(10)

	fmt.Println(NEXT, KEY, PREV)

	if FREE != -1 {
		t.Errorf("wanted -1, got %d", FREE)
	}
}
