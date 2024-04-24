package dsa

import "fmt"

type Item[T any] struct {
	Data T
	Next *Item[T]
}

type LinkedList[T any] struct {
	head *Item[T]
	tail *Item[T]
	size int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Add an item to the tail of the list
func (l *LinkedList[T]) Add(data T) {
	newItem := &Item[T]{
		Data: data,
		Next: nil,
	}

	if l.head == nil {
		l.head = newItem
		l.tail = newItem
	}

	l.tail.Next = newItem
	l.tail = newItem

	l.size++
}

func (l *LinkedList[T]) AddAt(index int, data T) {
	newItem := &Item[T]{
		Data: data,
		Next: nil,
	}

	if index >= l.size {
		l.tail.Next = newItem
		l.tail = newItem
	}

	currentItem := l.head

	for i := 1; i < index; i++ {
		currentItem = currentItem.Next
	}

	newItem.Next = currentItem.Next
	currentItem.Next = newItem

	l.size++
}

// Remove the last item
func (l *LinkedList[T]) Remove() {
	currentItem := l.head
	for i := 1; i < l.size-1; i++ {
		currentItem = currentItem.Next
	}

	currentItem.Next = nil
	l.size--
}

func (l *LinkedList[T]) RemoveAt(index int) {
	currentItem := l.head
	for i := 1; i < index; i++ {
		currentItem = currentItem.Next
	}

	currentItem.Next = currentItem.Next.Next
	l.size--
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) String() string {
	str := ""
	currentItem := l.head
	for currentItem.Next != nil {
		str += fmt.Sprintf("%v -> ", currentItem.Data)
	}
	str += "nil"
	return str
}
