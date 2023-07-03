package omdb

import (
	"errors"
	"testing"

	"github.com/krissukoco/omdb-api/internal/models"
)

func TestService(t *testing.T) {
	movies := []*models.Movie{
		{"batman", "Batman Begins", "2000", "R", "Action", "Batman", "Christopher Nolan", "English", "USA", "movie", "https://www.google.com", 0, 0, nil},
		{"shawshank", "Shawshank Redemption", "1970", "R", "Action", "Prison Escape", "Some Director", "English", "USA", "movie", "https://www.google.com", 0, 0, nil},
	}

	repo := NewMockRepository(movies)

	service := NewService(repo)

	// Get by id
	movie, err := service.GetMovieByID("batman")
	if err != nil {
		t.Fatalf("failed to get movie by id: %v", err)
	}
	if movie == nil {
		t.Fatal("movie is nil")
	}
	if movie.Id != "batman" {
		t.Fatalf("expected movie id to be batman, got %s", movie.Id)
	}

	// Get by non-existing id
	movie, err = service.GetMovieByID("non-existing")
	if !errors.Is(err, ErrMovieNotFound) {
		t.Fatalf("expected error to be ErrMovieNotFound, got %v", err)
	}
	if movie != nil {
		t.Fatal("movie is not nil")
	}

	// Search movies
	_, _, err = service.SearchMovies("Batman", "movie", 1)
	if err != nil {
		t.Fatalf("failed to search movies: %v", err)
	}
	// if count != 1 {
	// 	t.Fatalf("expected count to be 1, got %d", count)
	// }
	// if len(listMovies) != 1 {
	// 	t.Fatalf("expected list movies to be 1, got %d", len(listMovies))
	// }

	// Search movies with invalid query
	_, _, err = service.SearchMovies("ab", "movie", 1)
	if !errors.Is(err, ErrMinQueryLength) {
		t.Fatalf("expected error to be ErrMinQueryLength, got %v", err)
	}

	// Search movies with invalid page
	_, _, err = service.SearchMovies("Batman", "movie", -1)
	if !errors.Is(err, ErrInvalidPage) {
		t.Fatalf("expected error to be ErrInvalidPage, got %v", err)
	}

	// Search movies with invalid type
	_, _, err = service.SearchMovies("Batman", "invalid", 1)
	if !errors.Is(err, ErrInvalidMovieType) {
		t.Fatal("expected error to be ErrInvalidMovieType")
	}

}
