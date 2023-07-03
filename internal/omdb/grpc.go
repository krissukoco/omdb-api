package omdb

import (
	"context"
	"errors"
	"log"

	"github.com/krissukoco/omdb-api/internal/api"
	pb "github.com/krissukoco/omdb-api/proto/omdb"
	"google.golang.org/grpc/codes"
)

type Server struct {
	pb.UnimplementedOMDBServiceServer
	Service Service
}

func NewServer(s Service) *Server {
	return &Server{Service: s}
}

var _ pb.OMDBServiceServer = (*Server)(nil)

func (s *Server) GetMovieByID(ctx context.Context, in *pb.GetMovieByIDRequest) (*pb.GetMovieByIDResponse, error) {
	m, err := s.Service.GetMovieByID(in.Id)
	if err != nil {
		if errors.Is(err, ErrMovieNotFound) {
			return nil, api.GrpcError(codes.NotFound, api.CodeMovieNotFound, err.Error())
		}

		log.Println("internal error:", err)
		return nil, api.GrpcError(codes.Internal, api.CodeInternal, "internal error")
	}
	if m.Actors == nil {
		m.Actors = make([]string, 0)
	}
	return &pb.GetMovieByIDResponse{
		Id:        m.Id,
		Title:     m.Title,
		Year:      m.Year,
		Rated:     m.Rated,
		Genre:     m.Genre,
		Plot:      m.Plot,
		Director:  m.Director,
		Language:  m.Language,
		Country:   m.Country,
		Type:      m.Type,
		PosterUrl: m.PosterUrl,
		Actors:    m.Actors,
	}, nil
}

func (s *Server) SearchMovies(ctx context.Context, in *pb.SearchMoviesRequest) (*pb.SearchMoviesResponse, error) {
	res, count, err := s.Service.SearchMovies(in.Query, in.Type, int(in.Page))
	if err != nil {
		if errors.Is(err, ErrMinQueryLength) {
			return nil, api.GrpcError(codes.InvalidArgument, api.CodeInvalidQueryLength, err.Error())
		}

		log.Println("internal error:", err)
		return nil, api.GrpcError(codes.Internal, api.CodeInternal, "internal error")
	}
	movies := make([]*pb.MovieResult, 0)
	for _, m := range res {
		movies = append(movies, &pb.MovieResult{
			Id:        m.Id,
			Title:     m.Title,
			Year:      m.Year,
			Type:      m.Type,
			PosterUrl: m.PosterUrl,
		})
	}

	return &pb.SearchMoviesResponse{
		Movies:       movies,
		TotalResults: uint64(count),
	}, nil
}
