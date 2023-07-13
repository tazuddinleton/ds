package objectpointer

import "errors"

var (
	NEXT [9]int
	KEY  [9]int
	PREV [9]int
)

/*
Empty memory ready to be used. free = 0
[1, 2, 3, 4, 5, 6, 7, 8, -1]
[9, 0, 0, 0, 0, 0, 0, 0, 0]
[0, 0, 0, 0, 0, 0, 0, 0, 0]

*/

// Points to FREE head, where object can be inserted. Negative means out of space.
var FREE int = -1

// Holds the head pointer
var HEAD int = -1

// Wire up the free spaces by connecting next pointer
func resetFree() {
	FREE = 0
	for i := 2; i <= 10; i++ {
		NEXT[i-2] = (i % 10) - 1
	}
}

func allocateObject(val int) error {
	if FREE == -1 {
		return errors.New("out of space")
	}

	head := FREE
	FREE = NEXT[FREE]
	KEY[head] = val
	PREV[head] = HEAD
	HEAD = head
	return nil
}
