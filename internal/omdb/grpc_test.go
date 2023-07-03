package omdb

import (
	"context"
	"strconv"
	"strings"
	"testing"

	"github.com/krissukoco/omdb-api/internal/api"
	"github.com/krissukoco/omdb-api/internal/models"
	pb "github.com/krissukoco/omdb-api/proto/omdb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var movies = []*models.Movie{
	{"batman", "Batman Begins", "2000", "R", "Action", "Batman", "Christopher Nolan", "English", "USA", "movie", "https://www.google.com", 0, 0, nil},
	{"shawshank", "Shawshank Redemption", "1970", "R", "Action", "Prison Escape", "Some Director", "English", "USA", "movie", "https://www.google.com", 0, 0, nil},
}

func TestGrpc(t *testing.T) {
	server := NewServer(NewService(NewMockRepository(movies)))

	// Get movie by id
	movie, err := server.GetMovieByID(context.Background(), &pb.GetMovieByIDRequest{Id: "batman"})
	if err != nil {
		t.Fatalf("failed to get movie by id: %v", err)
	}
	if movie == nil {
		t.Fatal("movie is nil")
	}

	// Get movie by non-existing id
	movie, err = server.GetMovieByID(context.Background(), &pb.GetMovieByIDRequest{Id: "non-existing"})
	if err == nil {
		t.Fatal("expected error to be not nil")
	}
	if movie != nil {
		t.Fatal("movie is not nil")
	}
	// assert error codes
	st, ok := status.FromError(err)
	if !ok {
		t.Fatal("expected error to be grpc status error")
	}
	if st.Code() != codes.NotFound {
		t.Fatalf("expected error code to be NotFound, got %s", st.Code())
	}
	msg := st.Message()
	splitMsg := strings.Split(msg, "__")
	if len(splitMsg) != 2 {
		t.Fatalf("expected message to be split by __, got %s", msg)
	}
	intCode, err := strconv.Atoi(splitMsg[0])
	if err != nil {
		t.Fatalf("failed to convert string to int: %v", err)
	}
	if intCode != api.CodeMovieNotFound {
		t.Fatalf("expected error code to be %d, got %d", api.CodeMovieNotFound, intCode)
	}

	// Search movies
	movies, err := server.SearchMovies(context.Background(), &pb.SearchMoviesRequest{Query: "Batman", Type: "movie", Page: 1})
	if err != nil {
		t.Fatalf("failed to search movies: %v", err)
	}
	if movies == nil {
		t.Fatal("movies is nil")
	}
	// if len(movies.Movies) != 1 {
	// 	t.Fatalf("expected movies length to be 1, got %d", len(movies.Movies))
	// }
	// Expect movie batman to exist
	found := false
	for _, m := range movies.Movies {
		if m.Id == "batman" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("expected movie batman to exist")
	}
	// Expect count
	if movies.TotalResults != 2 {
		t.Fatalf("expected movies totalresult to be 1, got %d", movies.TotalResults)
	}

}
