package repository

import (
	"errors"
)

type Movie struct {
	Id    int
	Title string
	Year  int
}

// Declaramos los errores como variables locales para identificarlos

var (
	ErrInvalidMovie       = errors.New("invalid movie")
	ErrContraintViolation = errors.New("contraint violation")
)

type MovieMap struct {
	movies map[int]Movie
	lastId int
}

func (m *MovieMap) Save(movie *Movie) (err error) {
	// Check validation before save a new movie
	if movie.Title == "" {
		err = ErrInvalidMovie
		return
	}

	if movie.Year < 1900 && movie.Year > 2100 {
		err = ErrInvalidMovie
		return
	}

	for _, m := range m.movies {
		if m.Title == movie.Title {
			err = ErrContraintViolation
			return
		}
	}

	// Incre,emtamos el id
	(*m).lastId++
	(*movie).Id = (*m).lastId
	m.movies[(*movie).Id] = *movie

	return nil
}

func NewMovieMap() {
	return
}

// fmt.Errorf("%w limit %d", 20, 2) encadenamiento de errores

func main() {
	// Inicializamos el map con capacidad de 3

	// Creamos una movie
	/*
		movie:= Movie{
			Title: "Matrix",
			Year: 1998,
			Id: 123,
		}*/

}
