package models

type MovieActor struct {
	Id        int `gorm:"primaryKey"`
	MovieId   string
	Actor     string
	CreatedAt int64 `gorm:"autoCreateTime:milli"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
}
