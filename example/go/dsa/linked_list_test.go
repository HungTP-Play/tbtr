package dsa

import (
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList[int]()
	if ll.head != nil || ll.tail != nil || ll.size != 0 {
		t.Errorf("Expected empty linked list, got head: %v, tail: %v, size: %d", ll.head, ll.tail, ll.size)
	}
}

func TestAdd(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)

	expected := []int{1, 2, 3}
	current := ll.head
	for i, val := range expected {
		if current.Data != val {
			t.Errorf("Expected value %d at index %d, got %d", val, i, current.Data)
		}
		current = current.Next
	}

	if ll.size != len(expected) {
		t.Errorf("Expected size %d, got %d", len(expected), ll.size)
	}
}

func TestAddAt(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Add(1)
	ll.Add(3)
	ll.AddAt(1, 2)

	expected := []int{1, 2, 3}
	current := ll.head
	for i, val := range expected {
		if current.Data != val {
			t.Errorf("Expected value %d at index %d, got %d", val, i, current.Data)
		}
		current = current.Next
	}

	if ll.size != len(expected) {
		t.Errorf("Expected size %d, got %d", len(expected), ll.size)
	}
}

func TestRemove(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Remove()

	expected := []int{1, 2}
	current := ll.head
	for i, val := range expected {
		if current.Data != val {
			t.Errorf("Expected value %d at index %d, got %d", val, i, current.Data)
		}
		current = current.Next
	}

	if ll.size != len(expected) {
		t.Errorf("Expected size %d, got %d", len(expected), ll.size)
	}
}

func TestRemoveAt(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.RemoveAt(1)

	expected := []int{1, 3}
	current := ll.head
	for i, val := range expected {
		if current.Data != val {
			t.Errorf("Expected value %d at index %d, got %d", val, i, current.Data)
		}
		current = current.Next
	}

	if ll.size != len(expected) {
		t.Errorf("Expected size %d, got %d", len(expected), ll.size)
	}
}

func TestSize(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)

	if ll.Size() != 3 {
		t.Errorf("Expected size 3, got %d", ll.Size())
	}
}
