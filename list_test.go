package algorithms

import (
	"testing"
)

func TestListInsert(t *testing.T) {
	l1 := &list{1, nil}
	l0 := &list{0, l1}

	l2 := l0.insert(2)
	if l2.String() != "2-0-1-end" {
		t.Errorf("%s insert(2) = %s, want 2-%s", l0, l2, l0)
	}
}

func TestListGet(t *testing.T) {
	l2 := &list{2, nil}
	l1 := &list{1, l2}
	l0 := &list{0, l1}

	got := l0.get(2)
	if got != l2 {
		t.Errorf("%s get(2) = %s, want %s", l0, got, l2)
	}
}

func TestListRemove(t *testing.T) {
	l2 := &list{2, nil}
	l1 := &list{1, l2}
	l0 := &list{0, l1}
	want2 := &list{2, nil}
	want := &list{0, want2}

	got := l0.remove(1)
	if got.String() != want.String() {
		t.Errorf("%s remove(1) = %s, want %s", l0, got, want)
	}
}
