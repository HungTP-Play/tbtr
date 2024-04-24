package dsa

import "fmt"

type Array[T any] struct {
	array []T
}

func NewArray[T any](size int) *Array[T] {
	return &Array[T]{
		array: make([]T, size),
	}
}

func (a *Array[T]) Get(index int) T {
	return a.array[index]
}

func (a *Array[T]) Set(index int, val T) {
	a.array[index] = val
}

func (a *Array[T]) Size() int {
	return len(a.array)
}

func (a *Array[T]) InsertAt(index int, val T) {
	for i := a.Size() - 1; i > index; i-- {
		a.array[i] = a.array[i-1]
	}

	a.array[index] = val
}

func (a *Array[T]) DeleteAt(index int) {
	for i := index; i < a.Size()-1; i++ {
		a.array[i] = a.array[i+1]
	}
}

func (a *Array[T]) String() string {
	str := "["
	for _, i := range a.array {
		str += fmt.Sprintf("%v ", i)
	}
	str += "]"
	return str
}
