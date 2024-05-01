package dsa

import (
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList[int]()
	if list.Size != 0 {
		t.Errorf("Expected size to be 0, got %d", list.Size)
	}
	if list.head != nil {
		t.Errorf("Expected head to be nil, got %v", list.head)
	}
	if list.tail != nil {
		t.Errorf("Expected tail to be nil, got %v", list.tail)
	}
}

func TestAppend(t *testing.T) {
	list := NewLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	expected := []int{1, 2, 3}
	current := list.head
	for i, val := range expected {
		if current.Data != val {
			t.Errorf("Expected value at index %d to be %d, got %d", i, val, current.Data)
		}
		current = current.Next
	}
	if list.Size != len(expected) {
		t.Errorf("Expected size to be %d, got %d", len(expected), list.Size)
	}
}

func TestInsertAtLinkedList(t *testing.T) {
	list := NewLinkedList[int]()
	list.Append(1)
	list.Append(3)
	list.InsertAt(1, 2)
	expected := []int{1, 2, 3}
	current := list.head
	for i, val := range expected {
		if current.Data != val {
			t.Errorf("Expected value at index %d to be %d, got %d", i, val, current.Data)
		}
		current = current.Next
	}
	if list.Size != len(expected) {
		t.Errorf("Expected size to be %d, got %d", len(expected), list.Size)
	}
}

func TestDeleteAt(t *testing.T) {
	list := NewLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.DeleteAt(1)
	expected := []int{1, 3}
	current := list.head
	for i, val := range expected {
		if current.Data != val {
			t.Errorf("Expected value at index %d to be %d, got %d", i, val, current.Data)
		}
		current = current.Next
	}
	if list.Size != len(expected) {
		t.Errorf("Expected size to be %d, got %d", len(expected), list.Size)
	}
}

func TestStringListedList(t *testing.T) {
	list := NewLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	expected := "[ 1 -> 2 -> 3 -> 4 -> nil ]"
	if list.String() != expected {
		t.Errorf("Expected string representation to be %s, got %s", expected, list.String())
	}
}
