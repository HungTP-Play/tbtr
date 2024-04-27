package dsa

import (
	"fmt"
)

type Array[T any] struct {
	arr  []T
	cap  int // The number of elements that the Array can hold
	size int // The real, number of holding values
}

func NewArray[T any](initialSize int) *Array[T] {
	return &Array[T]{
		arr:  make([]T, initialSize),
		size: 0,
		cap:  initialSize,
	}
}

func (a *Array[T]) Set(index int, val T) {
	a.arr[index] = val
}

func (a *Array[T]) Get(index int) T {
	return a.arr[index]
}

func (a *Array[T]) resize() {
	newSize := a.cap * 2
	newArr := make([]T, newSize)
	copy(newArr, a.arr)
	a.arr = newArr
	a.cap = newSize
}

func (a *Array[T]) InsertAt(index int, val T) {
	if a.size >= a.cap {
		a.resize()
	}

	for i := a.cap - 1; i > index; i-- {
		a.arr[i] = a.arr[i-1]
	}

	a.arr[index] = val
	a.size++
}

func (a *Array[T]) RemoveAt(index int) {
	for i := index; i < a.cap-1; i++ {
		a.arr[i] = a.arr[i+1]
	}

	a.size--
}

func (a *Array[T]) Size() int {
	return a.size
}

func (a *Array[T]) Cap() int {
	return a.cap
}

func (a *Array[T]) String() string {
	str := "["
	for _, val := range a.arr {
		str += fmt.Sprintf("%v ", val)
	}

	str += "]"
	return str
}
