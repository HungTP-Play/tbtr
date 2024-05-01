package dsa

import (
	"testing"
)

func TestNewFilmRanking(t *testing.T) {
	fr := NewFilmRanking()
	if fr == nil {
		t.Errorf("NewFilmRanking() returned nil")
	}
	if len(fr.arr) != 0 {
		t.Errorf("Expected empty array, got length %d", len(fr.arr))
	}
}

func TestAdd(t *testing.T) {
	fr := NewFilmRanking()
	film1 := &Film{ID: "1", Name: "Film 1", Score: 8.5}
	film2 := &Film{ID: "2", Name: "Film 2", Score: 7.2}
	fr.Add(film1)
	fr.Add(film2)

	if len(fr.arr) != 2 {
		t.Errorf("Expected array length 2, got %d", len(fr.arr))
	}
	if fr.arr[0] != film1 || fr.arr[1] != film2 {
		t.Errorf("Incorrect films added to the array")
	}
}

func TestAvgRanking(t *testing.T) {
	fr := NewFilmRanking()
	film1 := &Film{ID: "1", Name: "Film 1", Score: 8.0}
	film2 := &Film{ID: "2", Name: "Film 2", Score: 7.0}
	film3 := &Film{ID: "3", Name: "Film 3", Score: 9.0}
	fr.Add(film1)
	fr.Add(film2)
	fr.Add(film3)

	avg := fr.AvgRanking()
	expectedAvg := float32(8.0)
	if avg != expectedAvg {
		t.Errorf("Expected average ranking %.2f, got %.2f", expectedAvg, avg)
	}
}

func TestHighest(t *testing.T) {
	fr := NewFilmRanking()
	film1 := &Film{ID: "1", Name: "Film 1", Score: 8.5}
	film2 := &Film{ID: "2", Name: "Film 2", Score: 7.2}
	film3 := &Film{ID: "3", Name: "Film 3", Score: 9.0}
	fr.Add(film1)
	fr.Add(film2)
	fr.Add(film3)

	highest := fr.Highest()
	if highest != film3 {
		t.Errorf("Expected highest ranking film to be %v, got %v", film3, highest)
	}
}

func TestLowest(t *testing.T) {
	fr := NewFilmRanking()
	film1 := &Film{ID: "1", Name: "Film 1", Score: 8.5}
	film2 := &Film{ID: "2", Name: "Film 2", Score: 7.2}
	film3 := &Film{ID: "3", Name: "Film 3", Score: 6.0}
	fr.Add(film1)
	fr.Add(film2)
	fr.Add(film3)

	lowest := fr.Lowest()
	if lowest != film3 {
		t.Errorf("Expected lowest ranking film to be %v, got %v", film3, lowest)
	}
}
