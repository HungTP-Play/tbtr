package dsa

type Film struct {
	ID    string
	Name  string
	Score float32 // Score ranking of the film, score in range [0.0, 10.0]
}

type FilmRanking struct {
	arr []*Film
}

func NewFilmRanking() *FilmRanking {
	return &FilmRanking{
		arr: make([]*Film, 0),
	}
}

func (f *FilmRanking) Add(film *Film) {
	f.arr = append(f.arr, film)
}

func (f *FilmRanking) AvgRanking() float32 {
	sum := float32(0.0)
	num := float32(0.0)
	for _, film := range f.arr {
		if film != nil {
			sum += film.Score
			num++
		}
	}

	return sum / num
}

func (f *FilmRanking) Highest() *Film {
	maxScore := float32(0.0)
	var highestRankingFilm *Film
	for _, film := range f.arr {
		if film != nil && film.Score > maxScore {
			maxScore = film.Score
			highestRankingFilm = film
		}
	}

	return highestRankingFilm
}

func (f *FilmRanking) Lowest() *Film {
	minScore := float32(11.0) // <- because ranking score is in the range [0.0, 10.0]
	var lowestRankingFilm *Film
	for _, val := range f.arr {
		if val != nil && val.Score < minScore {
			minScore = val.Score
			lowestRankingFilm = val
		}
	}

	return lowestRankingFilm
}
