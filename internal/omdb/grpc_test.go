package omdb

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/krissukoco/omdb-api/internal/models"
	"github.com/krissukoco/omdb-api/internal/tests"
	pb "github.com/krissukoco/omdb-api/proto/omdb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var movies = []*models.Movie{
	{"batman", "Batman Begins", "2000", "R", "Action", "Batman", "Christopher Nolan", "English", "USA", "movie", "https://www.google.com", 0, 0, nil},
	{"shawshank", "Shawshank Redemption", "1970", "R", "Action", "Prison Escape", "Some Director", "English", "USA", "movie", "https://www.google.com", 0, 0, nil},
}

func startGrpcServer() {
	lis := tests.NewTestGrpcListener()
	s := grpc.NewServer()
	pb.RegisterOMDBServiceServer(s, NewServer(NewService(NewMockRepository(movies))))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}
}

func bufdialer(lis *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, address string) (net.Conn, error) {
		return lis.Dial()
	}
}

func TestGrpc(t *testing.T) {
	go startGrpcServer()

	lis := tests.NewTestGrpcListener()
	defer lis.Close()

	// ctx := context.Background()
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(bufdialer(lis)), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}

	client := pb.NewOMDBServiceClient(conn)
	println(client)

	// Get by id
	ctx := context.Background()
	movie, err := client.GetMovieByID(ctx, &pb.GetMovieByIDRequest{Id: "batman"})
	if err != nil {
		t.Fatalf("failed to get movie by id: %v", err)
	}
	if movie == nil {
		t.Fatal("movie is nil")
	}

}
