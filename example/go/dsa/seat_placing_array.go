package dsa

import (
	"errors"
	"fmt"
)

const (
	NONE      = "NONE"
	PLACED    = "PLACED"
	AVAILABLE = "AVAILABLE"
)

type SeatBooking struct {
	seats [][]string // string can only by one of NONE, PLACED, or AVAILABLE
	Rows  int
	Cols  int
}

// Create new seat booking for a cine
//
// The initial map shoule contains only NONE, or AVIALBLE seats
func NewSeatBooking(initialCineMap [][]string) *SeatBooking {
	rows := len(initialCineMap)
	var cols int
	if rows == 0 {
		cols = 0
	} else {
		cols = len(initialCineMap[0])
	}
	return &SeatBooking{
		seats: initialCineMap,
		Rows:  rows,
		Cols:  cols,
	}
}

func (s *SeatBooking) Check(row, col int) bool {
	return s.seats[row][col] == AVAILABLE
}

func (s *SeatBooking) Book(row, col int) (bool, error) {
	if row < 0 || row >= s.Rows {
		return false, fmt.Errorf("index out of range, your row index=%d", row)
	}

	if col < 0 || col >= s.Cols {
		return false, fmt.Errorf("index out of range, your col index=%d", col)
	}

	if !s.Check(row, col) {
		return false, errors.New("the seat is already booked")
	}

	s.seats[row][col] = PLACED
	return true, nil
}

func (s *SeatBooking) String() string {
	str := ""
	if len(s.seats) == 0 {
		return str
	}

	for i := 0; i < s.Rows; i++ {
		for j := 0; j < s.Cols; j++ {
			if s.seats[i][j] == NONE {
				str += " "
				continue
			}

			if s.seats[i][j] == AVAILABLE {
				str += "O"
				continue
			}

			str += "X"
		}

		str += "\n"
	}

	return str
}
