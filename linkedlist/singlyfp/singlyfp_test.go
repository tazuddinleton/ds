package singlyfp

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	l := New(1, New(2, nil))
	if l.data != 1 {
		t.Errorf("wanted 1 got %d", l.data)
	}
}

func TestNewList(t *testing.T) {
	l := NewList("Hello", "World", "!")
	if l.data != "Hello" {
		t.Errorf("wanted \"Hello\", got %s", l.data)
	}
}

func TestFlatten(t *testing.T) {
	l := NewList("Hello", "World", "!")
	fl := Flatten(l)
	if len(fl) != 3 {
		t.Errorf("wanted len 3, got %d", len(fl))
	}
	cont := strings.Join(fl, " ")
	if cont != "Hello World !" {
		t.Errorf("wanted \"Hello World !\", got %s", cont)
	}
}

func TestFlattenSelect(t *testing.T) {

	type Duck struct {
		Name string
	}
	l := NewList("Hello", "World", "!")
	fl := FlattenSelect(l, func(t string) string { return t })
	if len(fl) != 3 {
		t.Errorf("wanted len 3, got %d", len(fl))
	}
	cont := strings.Join(fl, " ")
	if cont != "Hello World !" {
		t.Errorf("wanted \"Hello World !\", got %s", cont)
	}

	l2 := NewList(&Duck{"Donald"}, &Duck{"Duck"})
	fl2 := FlattenSelect(l2, func(t *Duck) string { return t.Name })

	if len(fl2) != 2 {
		t.Errorf("wanted len 2, got %d", len(fl2))
	}

	cont2 := strings.Join(fl2, " ")

	if cont2 != "Donald Duck" {
		t.Errorf("wanted \"Donald Duck\", got %s", cont2)
	}
}

func TestInsertHead(t *testing.T) {
	l := NewList(1, 2, 3, 4)
	l = InsertHead(55, l)

	if l.data != 55 {
		t.Errorf("wanted 55, got %d", l.data)
	}
}

func TestLast(t *testing.T) {
	l := NewList(1, 2, 3)
	lst := 4
	l = Append(l, lst)

	got := Last(l)

	if got != lst {
		t.Errorf("wanted %d, got %d", lst, got)
	}
}

func TestLastWhenNil(t *testing.T) {
	var lst int
	l := Append(nil, lst)

	got := Last(l)

	if got != lst {
		t.Errorf("wanted %d, got %d", lst, got)
	}
}

func TestAppend(t *testing.T) {
	l := NewList(1, 2, 3)
	lst := 4
	l = Append(l, lst)

	got := Last(l)

	if got != lst && l.data != 1 {
		t.Errorf("wanted %d, got %d", lst, got)
	}

	if l.data != 1 {
		t.Errorf("wanted %d, got %d", 1, l.data)
	}

}

func TestRevers(t *testing.T) {
	l := NewList(1, 2, 3, 4, 5)
	l = Reverse(l)
	if l.data != 5 {
		t.Errorf("wanted 5, got %d", l.data)
	}
}
