package dsa

import "fmt"

type Node[T any] struct {
	Data T
	Next *Node[T]
	Prev *Node[T]
}

type LinkedList[T any] struct {
	Size int
	head *Node[T]
	tail *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{Size: 0}
}

func (l *LinkedList[T]) Append(data T) {
	newNode := &Node[T]{
		Data: data,
	}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		if l.tail.Prev == nil {
			l.head.Next = newNode
			newNode.Prev = l.head
			l.tail = newNode
		} else {
			l.tail.Prev.Next.Next = newNode
			newNode.Prev = l.tail
			l.tail = newNode
		}
	}

	l.Size++
}

func (l *LinkedList[T]) InsertAt(index int, data T) {
	fromStart := index <= (l.Size / 2)
	newNode := &Node[T]{
		Data: data,
	}

	if fromStart {
		currentNode := l.head
		for i := 1; i < index; i++ {
			currentNode = currentNode.Next
		}

		newNode.Next = currentNode.Next
		newNode.Prev = currentNode

		currentNode.Next = newNode
	} else {
		currentNode := l.tail
		for i := l.Size - 2; i >= index; i-- {
			currentNode = currentNode.Prev
		}

		newNode.Next = currentNode
		newNode.Prev = currentNode.Prev

		currentNode.Prev.Next = newNode
	}

	l.Size++
}

func (l *LinkedList[T]) DeleteAt(index int) {
	fromStart := index <= (l.Size / 2)

	if fromStart {
		currentNode := l.head
		for i := 1; i < index; i++ {
			currentNode = currentNode.Next
		}

		currentNode.Next = currentNode.Next.Next
		currentNode.Next.Prev = currentNode
	} else {
		currentNode := l.tail
		for i := l.Size - 2; i >= index; i-- {
			currentNode = currentNode.Prev
		}

		currentNode.Prev.Next = currentNode.Next
		currentNode.Prev.Next.Prev = currentNode.Prev
	}

	l.Size--
}

func (l *LinkedList[T]) String() string {
	str := "[ "
	currentNode := l.head
	for currentNode.Next != nil {
		str += fmt.Sprintf("%v -> ", currentNode.Data)
		currentNode = currentNode.Next
	}
	str += fmt.Sprintf("%v -> ", currentNode.Data)
	str += "nil ]"
	return str
}
