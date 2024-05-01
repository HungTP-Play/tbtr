package dsa

import (
	"reflect"
	"testing"
)

func TestNewInventory(t *testing.T) {
	warningQuan := 10
	inv := NewInventory(warningQuan)

	if inv.WarningQuan != warningQuan {
		t.Errorf("NewInventory() failed to initialize WarningQuan correctly")
	}

	if len(inv.products) != 0 {
		t.Errorf("NewInventory() failed to initialize products slice correctly")
	}
}

func TestInventoryAdd(t *testing.T) {
	inv := NewInventory(10)
	p1 := &Product{ID: "1", Name: "Product 1", Quantity: 20}
	p2 := &Product{ID: "2", Name: "Product 2", Quantity: 30}

	inv.Add(p1)
	inv.Add(p2)

	if len(inv.products) != 2 {
		t.Errorf("Add() failed to add products correctly")
	}

	if !reflect.DeepEqual(inv.products[0], p1) || !reflect.DeepEqual(inv.products[1], p2) {
		t.Errorf("Add() failed to add products correctly")
	}
}

func TestWarningProducts(t *testing.T) {
	inv := NewInventory(20)
	p1 := &Product{ID: "1", Name: "Product 1", Quantity: 10}
	p2 := &Product{ID: "2", Name: "Product 2", Quantity: 30}
	p3 := &Product{ID: "3", Name: "Product 3", Quantity: 15}

	inv.Add(p1)
	inv.Add(p2)
	inv.Add(p3)

	warningProducts := inv.WarningProducts()

	if len(warningProducts) != 2 {
		t.Errorf("WarningProducts() failed to return correct number of warning products")
	}

	if !reflect.DeepEqual(warningProducts[0], p1) || !reflect.DeepEqual(warningProducts[1], p3) {
		t.Errorf("WarningProducts() failed to return correct warning products")
	}
}

func TestUpdateQuan(t *testing.T) {
	inv := NewInventory(10)
	p1 := &Product{ID: "1", Name: "Product 1", Quantity: 20}
	p2 := &Product{ID: "2", Name: "Product 2", Quantity: 30}

	inv.Add(p1)
	inv.Add(p2)

	ok, err := inv.UpdateQuan("1", 25)
	if !ok || err != nil {
		t.Errorf("UpdateQuan() failed to update quantity for existing product")
	}

	if p1.Quantity != 25 {
		t.Errorf("UpdateQuan() failed to update quantity correctly")
	}

	ok, err = inv.UpdateQuan("3", 40)
	if ok || err == nil {
		t.Errorf("UpdateQuan() failed to handle non-existing product correctly")
	}
}

func TestFindId(t *testing.T) {
	inv := NewInventory(10)
	p1 := &Product{ID: "1", Name: "Product 1", Quantity: 20}
	p2 := &Product{ID: "2", Name: "Product 2", Quantity: 30}

	inv.Add(p1)
	inv.Add(p2)

	foundProduct := inv.FindId("1")
	if !reflect.DeepEqual(foundProduct, p1) {
		t.Errorf("FindId() failed to find product by ID")
	}

	foundProduct = inv.FindId("3")
	if foundProduct != nil {
		t.Errorf("FindId() failed to handle non-existing product ID correctly")
	}
}
func TestFindName(t *testing.T) {
	inv := NewInventory(10)
	p1 := &Product{ID: "1", Name: "Product 1", Quantity: 20}
	p2 := &Product{ID: "2", Name: "Product 2", Quantity: 30}

	inv.Add(p1)
	inv.Add(p2)

	foundProduct := inv.FindName("Product 1")
	if !reflect.DeepEqual(foundProduct, p1) {
		t.Errorf("FindName() failed to find product by name")
	}

	foundProduct = inv.FindName("Non-existing Product")
	if foundProduct != nil {
		t.Errorf("FindName() failed to handle non-existing product name correctly")
	}
}
