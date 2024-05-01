package dsa

import (
	"reflect"
	"testing"
)

func TestNewSeatBooking(t *testing.T) {
	initialCineMap := [][]string{
		{"NONE", "AVAILABLE", "NONE"},
		{"AVAILABLE", "NONE", "AVAILABLE"},
		{"NONE", "AVAILABLE", "NONE"},
	}

	sb := NewSeatBooking(initialCineMap)

	if !reflect.DeepEqual(sb.seats, initialCineMap) {
		t.Errorf("NewSeatBooking() failed to initialize seats correctly")
	}
}

func TestCheck(t *testing.T) {
	initialCineMap := [][]string{
		{"NONE", "AVAILABLE", "NONE"},
		{"AVAILABLE", "PLACED", "AVAILABLE"},
		{"NONE", "AVAILABLE", "NONE"},
	}

	sb := NewSeatBooking(initialCineMap)

	if sb.Check(0, 0) {
		t.Errorf("Check() failed for seat (0, 0)")
	}

	if !sb.Check(0, 1) {
		t.Errorf("Check() failed for seat (0, 1)")
	}

	if sb.Check(1, 1) {
		t.Errorf("Check() failed for seat (1, 1)")
	}
}

func TestBook(t *testing.T) {
	initialCineMap := [][]string{
		{"NONE", "AVAILABLE", "NONE"},
		{"AVAILABLE", "NONE", "AVAILABLE"},
		{"NONE", "AVAILABLE", "NONE"},
	}

	sb := NewSeatBooking(initialCineMap)

	// Book a valid seat
	ok, err := sb.Book(0, 1)
	if !ok || err != nil {
		t.Errorf("Book() failed for seat (0, 1)")
	}

	// Try to book an already booked seat
	ok, err = sb.Book(0, 1)
	if ok || err == nil {
		t.Errorf("Book() failed to handle already booked seat (0, 1)")
	}

	// Try to book an invalid seat
	ok, err = sb.Book(3, 0)
	if ok || err == nil {
		t.Errorf("Book() failed to handle invalid seat (3, 0)")
	}
}

func TestSeatBookingString(t *testing.T) {
	initialCineMap := [][]string{
		{"NONE", "AVAILABLE", "NONE"},
		{"AVAILABLE", "PLACED", "AVAILABLE"},
		{"NONE", "AVAILABLE", "NONE"},
	}

	sb := NewSeatBooking(initialCineMap)

	expected := " O \nOXO\n O \n"
	if sb.String() != expected {
		t.Errorf("String() failed to generate correct output")
	}
}
