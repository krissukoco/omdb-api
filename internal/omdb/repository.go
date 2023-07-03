package omdb

import (
	"errors"

	"github.com/krissukoco/omdb-api/internal/models"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

var (
	ErrMovieNotFound = errors.New("movie not found")
)

type Movie struct {
	Id        string
	Title     string
	Year      string
	Rated     string
	Genre     string
	Plot      string
	Director  string
	Language  string
	Country   string
	Type      string
	PosterUrl string
	CreatedAt int64
	UpdatedAt int64
	Actors    []string
}

type MovieSimple struct {
	Id        string
	Title     string
	Year      string
	PosterUrl string
	Type      string
}

type Repository interface {
	GetMovieByID(id string) (*models.Movie, error)
	SearchMovies(query string, movieType string, page int) ([]*MovieSimple, int64, error)
	InsertMovie(m *models.Movie, actors []string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&models.Movie{})
	db.AutoMigrate(&models.MovieActor{})

	return &repository{db}
}

func NewMovieId() string {
	return ulid.Make().String()
}

func (r *repository) GetMovieByID(id string) (*models.Movie, error) {
	var movie models.Movie
	if err := r.db.Model(&models.Movie{}).Where("id = ?", id).Preload("Actors").First(&movie).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMovieNotFound
		}
		return nil, err
	}

	return &movie, nil
}

func (r *repository) SearchMovies(query string, movieType string, page int) ([]*MovieSimple, int64, error) {
	var movies []*MovieSimple
	var count int64
	// raw := "SELECT id, title, year, type FROM movies WHERE title LIKE ?"
	q := r.db.Table("movies").Where("title LIKE ?", "%"+query+"%")
	if movieType != "" {
		q = q.Where("type = ?", movieType)
	}
	q.Count(&count)
	q = q.Select("id, title, year, type, poster_url").Offset((page - 1) * 10).Limit(10)
	rows, err := q.Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var movie MovieSimple
		if err := r.db.ScanRows(rows, &movie); err != nil {
			return nil, 0, err
		}
		movies = append(movies, &movie)
	}
	return movies, count, nil
}

func (r *repository) InsertMovie(m *models.Movie, actors []string) error {
	// m.Id = NewMovieId()

	if err := r.db.Create(m).Error; err != nil {
		return err
	}

	// // Create actors
	// for _, a := range m.Actors {
	// 	a.MovieId = m.Id
	// 	if err := r.db.Create(a).Error; err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

type mockrepository struct {
	items []*models.Movie
}

func NewMockRepository(items []*models.Movie) Repository {
	return &mockrepository{items}
}

func (m *mockrepository) GetMovieByID(id string) (*models.Movie, error) {
	for _, item := range m.items {
		if item.Id == id {
			return item, nil
		}
	}
	return nil, ErrMovieNotFound
}

func (m *mockrepository) SearchMovies(query string, movieType string, page int) ([]*MovieSimple, int64, error) {
	movies := make([]*MovieSimple, 0)
	for _, item := range m.items {
		if item.Type == movieType {
			movies = append(movies, &MovieSimple{
				Id:        item.Id,
				Title:     item.Title,
				Year:      item.Year,
				PosterUrl: item.PosterUrl,
				Type:      item.Type,
			})
		}
	}
	return movies, int64(len(movies)), nil
}

func (m *mockrepository) InsertMovie(movie *models.Movie, actors []string) error {
	m.items = append(m.items, movie)
	return nil
}
