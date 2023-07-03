package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DefaultPostgresPort     = 5432
	DefaultPostgresTimezone = "Asia/Jakarta"
)

func postgresDsn(host, user, password, dbname, timezone string, port int, enableSsl bool) string {
	sslMode := "disable"
	if enableSsl {
		sslMode = "enable"
	}
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s port=%d TimeZone=%s", host, user, password, dbname, sslMode, port, timezone)
}

func NewPostgresGorm(host, user, password, dbname, timezone string, port int, enableSsl bool) (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.Open(postgresDsn(host, user, password, dbname, timezone, port, enableSsl)),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
