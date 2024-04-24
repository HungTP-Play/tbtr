package dsa

import (
	"reflect"
	"testing"
)

func TestNewArray(t *testing.T) {
	arr := NewArray[int](5)
	if arr.Size() != 5 {
		t.Errorf("Expected size 5, got %d", arr.Size())
	}
}

func TestGet(t *testing.T) {
	arr := NewArray[int](5)
	arr.Set(2, 10)
	if arr.Get(2) != 10 {
		t.Errorf("Expected value at index 2 to be 10, got %d", arr.Get(2))
	}
}

func TestSet(t *testing.T) {
	arr := NewArray[string](3)
	arr.Set(1, "hello")
	if arr.Get(1) != "hello" {
		t.Errorf("Expected value at index 1 to be 'hello', got '%s'", arr.Get(1))
	}
}

func TestInsertAt(t *testing.T) {
	arr := NewArray[int](5)
	for i := 0; i < 5; i++ {
		arr.Set(i, i)
	}
	arr.InsertAt(2, 10)
	expected := []int{0, 1, 10, 2, 3}
	if !reflect.DeepEqual(arr.array, expected) {
		t.Errorf("Expected array %v, got %v", expected, arr.array)
	}
}

func TestDeleteAt(t *testing.T) {
	arr := NewArray[int](5)
	for i := 0; i < 5; i++ {
		arr.Set(i, i)
	}
	arr.DeleteAt(2)
	expected := []int{0, 1, 3, 4, 4}
	if !reflect.DeepEqual(arr.array, expected) {
		t.Errorf("Expected array %v, got %v", expected, arr.array)
	}
}

func TestString(t *testing.T) {
	arr := NewArray[int](3)
	arr.Set(0, 1)
	arr.Set(1, 2)
	arr.Set(2, 3)
	expected := "[1 2 3 ]"
	if arr.String() != expected {
		t.Errorf("Expected string '%s', got '%s'", expected, arr.String())
	}
}
