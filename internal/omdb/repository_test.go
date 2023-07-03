package omdb

import (
	"testing"

	"github.com/krissukoco/omdb-api/internal/models"
	"github.com/krissukoco/omdb-api/internal/tests"
)

func TestRepository_Read(t *testing.T) {
	db, err := tests.NewTestDb()
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	db.Exec("DELETE FROM movies")
	db.Exec("DELETE FROM movie_actors")

	repo := NewRepository(db)
	if repo == nil {
		t.Fatal("repository is nil")
	}

	// Create fake movies
	movies := []*models.Movie{
		{"batman", "Batman Begins", "2000", "R", "Action", "Batman", "Christopher Nolan", "English", "USA", "movie", "https://www.google.com", 0, 0, nil},
		{"shawshank", "Shawshank Redemption", "1970", "R", "Action", "Prison Escape", "Some Director", "English", "USA", "movie", "https://www.google.com", 0, 0, nil},
	}
	for _, m := range movies {
		if err := repo.InsertMovie(m, []string{"Christian Bale", "Michael Caine"}); err != nil {
			t.Fatalf("failed to insert movie: %v", err)
		}
	}

	// Get by id
	movie, err := repo.GetMovieByID("batman")
	if err != nil {
		t.Fatalf("failed to get movie by id: %v", err)
	}
	if movie == nil {
		t.Fatal("movie is nil")
	}
	if movie.Id != "batman" {
		t.Fatalf("expected movie id to be batman, got %s", movie.Id)
	}

	// Search movies
	listMovies, count, err := repo.SearchMovies("Batman", "movie", 1)
	if err != nil {
		t.Fatalf("failed to search movies: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected count to be 1, got %d", count)
	}
	if len(listMovies) != 1 {
		t.Fatalf("expected list movies to be 1, got %d", len(listMovies))
	}

}
