package algorithms

import (
	"fmt"
	"strings"
)

// List is a singly-linked list
type list struct {
	item int
	next *list
}

func (l *list) get(target int) *list {
	if l == nil {
		return nil
	}
	if l.item == target {
		return l
	}
	return l.next.get(target)
}

// Insert simply adds the new item to the head of the list
// Caller must assign the returned *list to replace the original
func (l *list) insert(target int) *list {
	// todo (zdjones) can I copy the address AND
	// replace the value in the same go?
	// I thinks it should evaluate the RH expresion first,
	// which will copy the address of the original list head.
	// Then the new list head will overwrite that original.
	return &list{
		item: target,
		next: l,
	}
}

// getPred finds the predecessor of the target item. Returns nil
// if the list is empty or the item is not found.
func (l *list) getPred(target int) *list {
	// Target is first item, or not in list (has no pred)
	if l.next == nil {
		return nil
	}
	// Found the target's predecessor
	if l.next.item == target {
		return l
	}
	// Check the rest of the list
	return l.next.getPred(target)
}

// remove deletes the first instance of the given item
// from the list.
func (l *list) remove(target int) *list {
	// Failed to find target
	if l.get(target) == nil {
		return l
	}
	// Target is first node = has no predecessor.
	if l.item == target {
		return l.next
	}
	// Target exists and is not first.
	// Find the target's predecessor
	pred := l.getPred(target)
	// Point the predecessor to the target's successor
	pred.next = pred.next.next
	return l
}

func (l *list) String() string {
	s := strings.Builder{}
	for l != nil {
		s.Write([]byte(fmt.Sprintf("%v-", l.item)))
		l = l.next
	}
	s.Write([]byte("end"))
	return s.String()
}
