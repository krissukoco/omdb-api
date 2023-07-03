package tests

import (
	"github.com/krissukoco/omdb-api/config"
	"github.com/krissukoco/omdb-api/internal/database"
	"gorm.io/gorm"
)

func NewTestDb() (*gorm.DB, error) {
	cfg, err := config.Load("test")
	if err != nil {
		return nil, err
	}
	dbc := cfg.Database
	db, err := database.NewPostgresGorm(dbc.Host, dbc.User, dbc.Password, dbc.Dbname, "Asia/Jakarta", dbc.Port, dbc.EnableSsl)
	if err != nil {
		return nil, err
	}
	return db, nil
}
