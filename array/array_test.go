package array

import "testing"

func TestReverse(t *testing.T) {
	arr := ArrayOf(1, 2, 3, 4)
	f := First(arr)
	Reverse(arr)
	l := Last(arr)

	if f != l {
		t.Errorf("wanted %d, got %d", f, l)
	}
}

func TestRotate(t *testing.T) {
	arr := ArrayOf(1, 2, 3, 4, 5)
	n := 2
	r := Rotate(arr, n)

	if len(arr) != len(r) {
		t.Errorf("got less elements after rotating array")
	}

	if First(r) != 3 {
		t.Errorf("wanted 2, got %d", First(r))
	}
	if Last(r) != 2 {
		t.Errorf("wanted 1, got %d", Last(r))
	}
}
