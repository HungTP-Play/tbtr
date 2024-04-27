package dsa

import (
	"reflect"
	"testing"
)

func TestNewArray(t *testing.T) {
	arr := NewArray[int](5)
	if arr.size != 0 {
		t.Errorf("Expected size to be 0, got %d", arr.size)
	}
	if arr.cap != 5 {
		t.Errorf("Expected capacity to be 5, got %d", arr.cap)
	}
}

func TestSet(t *testing.T) {
	arr := NewArray[int](5)
	arr.Set(0, 10)
	if arr.Get(0) != 10 {
		t.Errorf("Expected value at index 0 to be 10, got %d", arr.Get(0))
	}
}

func TestInsertAt(t *testing.T) {
	arr := NewArray[int](2)
	arr.InsertAt(0, 10)
	arr.InsertAt(1, 20)
	arr.InsertAt(1, 30)
	expected := []int{10, 30, 20}
	if !reflect.DeepEqual(arr.arr[:arr.size], expected) {
		t.Errorf("Expected array to be %v, got %v", expected, arr.arr[:arr.size])
	}
}

func TestRemoveAtArray(t *testing.T) {
	arr := NewArray[int](5)
	arr.InsertAt(0, 10)
	arr.InsertAt(1, 20)
	arr.InsertAt(2, 30)
	arr.RemoveAt(1)
	expected := []int{10, 30}
	if !reflect.DeepEqual(arr.arr[:arr.size], expected) {
		t.Errorf("Expected array to be %v, got %v", expected, arr.arr[:arr.size])
	}
}

func TestArraySize(t *testing.T) {
	arr := NewArray[int](5)
	arr.InsertAt(0, 10)
	arr.InsertAt(1, 20)
	if arr.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", arr.Size())
	}
}

func TestCap(t *testing.T) {
	arr := NewArray[int](5)
	if arr.Cap() != 5 {
		t.Errorf("Expected capacity to be 5, got %d", arr.Cap())
	}
}

func TestString(t *testing.T) {
	arr := NewArray[int](5)
	arr.InsertAt(0, 10)
	arr.InsertAt(1, 20)
	expected := "[10 20 0 0 0 ]"
	if arr.String() != expected {
		t.Errorf("Expected string representation to be %s, got %s", expected, arr.String())
	}
}
