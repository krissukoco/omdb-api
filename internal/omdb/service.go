package omdb

import (
	"errors"
	"strings"
)

var (
	ErrMinQueryLength   = errors.New("query length must be at least 3 characters")
	ErrInvalidPage      = errors.New("page must be greater than or equal to 1")
	ErrInvalidMovieType = errors.New("movie type must be one of movie, series, or episode")
)

var (
	MovieTypes = []string{"movie", "series", "episode"}
)

type Service interface {
	GetMovieByID(id string) (*Movie, error)
	SearchMovies(query string, movieType string, page int) ([]*MovieSimple, int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetMovieByID(id string) (*Movie, error) {
	m, err := s.repo.GetMovieByID(id)
	if err != nil {
		return nil, err
	}
	actors := make([]string, 0)
	for _, a := range m.Actors {
		actors = append(actors, a.Actor)
	}
	return &Movie{
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
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		Actors:    actors,
	}, nil
}

func (s *service) SearchMovies(query string, movieType string, page int) ([]*MovieSimple, int64, error) {
	if len(query) < 3 {
		return nil, 0, ErrMinQueryLength
	}
	if page < 1 {
		return nil, 0, ErrInvalidPage
	}
	if page == 0 {
		page = 1
	}
	mtype := ""
	for _, t := range MovieTypes {
		if t == strings.ToLower(movieType) {
			mtype = t
			break
		}
	}
	if mtype == "" {
		return nil, 0, ErrInvalidMovieType
	}

	m, count, err := s.repo.SearchMovies(query, mtype, page)
	if err != nil {
		return nil, 0, err
	}
	return m, count, nil
}
