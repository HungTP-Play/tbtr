package dsa

import (
	"reflect"
	"testing"
)

func TestNewCart(t *testing.T) {
	cart := NewCart()
	if len(cart.arr) != 0 {
		t.Errorf("Expected cart to be empty, got length %d", len(cart.arr))
	}
}

func TestCartGet(t *testing.T) {
	cart := NewCart()
	p1 := &Product{ID: "1", Name: "Product 1", Price: 10.0, Quantity: 2}
	p2 := &Product{ID: "2", Name: "Product 2", Price: 20.0, Quantity: 1}
	cart.Add(p1)
	cart.Add(p2)

	if !reflect.DeepEqual(cart.Get(0), *p1) {
		t.Errorf("Expected %v, got %v", *p1, cart.Get(0))
	}
	if !reflect.DeepEqual(cart.Get(1), *p2) {
		t.Errorf("Expected %v, got %v", *p2, cart.Get(1))
	}
}

func TestCartCartSet(t *testing.T) {
	cart := NewCart()
	p1 := &Product{ID: "1", Name: "Product 1", Price: 10.0, Quantity: 2}
	p2 := &Product{ID: "2", Name: "Product 2", Price: 20.0, Quantity: 1}
	cart.Add(p1)
	cart.Add(p2)

	p3 := &Product{ID: "3", Name: "Product 3", Price: 30.0, Quantity: 3}
	cart.Set(1, p3)

	if !reflect.DeepEqual(cart.Get(1), *p3) {
		t.Errorf("Expected %v, got %v", *p3, cart.Get(1))
	}
}

func TestCartAdd(t *testing.T) {
	cart := NewCart()
	p1 := &Product{ID: "1", Name: "Product 1", Price: 10.0, Quantity: 2}
	p2 := &Product{ID: "2", Name: "Product 2", Price: 20.0, Quantity: 1}
	cart.Add(p1)
	cart.Add(p2)

	if len(cart.arr) != 2 {
		t.Errorf("Expected cart length to be 2, got %d", len(cart.arr))
	}
	if !reflect.DeepEqual(cart.Get(0), *p1) {
		t.Errorf("Expected %v, got %v", *p1, cart.Get(0))
	}
	if !reflect.DeepEqual(cart.Get(1), *p2) {
		t.Errorf("Expected %v, got %v", *p2, cart.Get(1))
	}
}

func TestCartDelete(t *testing.T) {
	cart := NewCart()
	p1 := &Product{ID: "1", Name: "Product 1", Price: 10.0, Quantity: 2}
	p2 := &Product{ID: "2", Name: "Product 2", Price: 20.0, Quantity: 1}
	p3 := &Product{ID: "3", Name: "Product 3", Price: 30.0, Quantity: 3}
	cart.Add(p1)
	cart.Add(p2)
	cart.Add(p3)

	cart.Delete(1)
	if !reflect.DeepEqual(cart.Get(0), *p1) || !reflect.DeepEqual(cart.Get(1), *p3) {
		t.Errorf("Unexpected cart contents: %v", cart.arr)
	}
}

func TestCartFilter(t *testing.T) {
	cart := NewCart()
	p1 := &Product{ID: "1", Name: "Product 1", Price: 10.0, Quantity: 2}
	p2 := &Product{ID: "2", Name: "Product 2", Price: 20.0, Quantity: 1}
	p3 := &Product{ID: "3", Name: "Product 3", Price: 30.0, Quantity: 3}
	cart.Add(p1)
	cart.Add(p2)
	cart.Add(p3)

	filtered := cart.Filter(func(p *Product, index int) bool {
		return p.Price > 15.0
	})

	if len(filtered) != 2 {
		t.Errorf("Expected filtered length to be 2, got %d", len(filtered))
	}
	if !reflect.DeepEqual(*filtered[0], *p2) || !reflect.DeepEqual(*filtered[1], *p3) {
		t.Errorf("Unexpected filtered products: %v", filtered)
	}
}

func TestCartTotal(t *testing.T) {
	cart := NewCart()
	p1 := &Product{ID: "1", Name: "Product 1", Price: 10.0, Quantity: 2}
	p2 := &Product{ID: "2", Name: "Product 2", Price: 20.0, Quantity: 1}
	p3 := &Product{ID: "3", Name: "Product 3", Price: 30.0, Quantity: 3}
	cart.Add(p1)
	cart.Add(p2)
	cart.Add(p3)

	total := cart.Total()

	expectedTotal := (float64(p1.Price) * float64(p1.Quantity)) + (float64(p2.Price) * float64(p2.Quantity)) + (float64(p3.Price) * float64(p3.Quantity))
	if total != expectedTotal {
		t.Errorf("Expected total to be %.2f, got %.2f", expectedTotal, total)
	}
}
