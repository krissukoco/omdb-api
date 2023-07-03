package models

type Movie struct {
	Id        string `gorm:"primaryKey"`
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
	CreatedAt int64        `gorm:"autoCreateTime:milli"`
	UpdatedAt int64        `gorm:"autoUpdateTime:milli"`
	Actors    []MovieActor `gorm:"foreignKey:MovieId"`
}
