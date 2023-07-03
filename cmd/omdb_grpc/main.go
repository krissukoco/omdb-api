package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/krissukoco/omdb-api/config"
	"github.com/krissukoco/omdb-api/internal/database"
	"github.com/krissukoco/omdb-api/internal/omdb"
	omdbPb "github.com/krissukoco/omdb-api/proto/omdb"
	"google.golang.org/grpc"
)

func getEnvironment() string {
	env, exists := os.LookupEnv("ENV")
	if exists {
		return env
	}

	return "local"
}

func main() {
	cfg, err := config.Load(getEnvironment())
	if err != nil {
		panic(err)
	}

	dbc := cfg.Database
	db, err := database.NewPostgresGorm(dbc.Host, dbc.User, dbc.Password, dbc.Dbname, "Asia/Jakarta", dbc.Port, dbc.EnableSsl)
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	srv := omdb.NewServer(omdb.NewService(omdb.NewRepository(db)))

	omdbPb.RegisterOMDBServiceServer(s, srv)

	log.Println("Starting server on port", cfg.Port)
	log.Fatal(s.Serve(listener))
}
