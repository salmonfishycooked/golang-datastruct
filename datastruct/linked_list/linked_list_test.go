package datastruct

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := NewLinkedList()
	list.insert(0, 2)
	list.insert(0, 5)
	list.insert(0, 6)
	list.print()
}
