# OMDB gRPC API
### By Kris Sukoco
### As a requirement for technical test completion on S16 Ventures. Written in Go/Golang.
<br/>

## How the code is structured
1. Entrypoint of the server is on `cmd` directory: `cmd/omdb_grpc/main.go`
2. `config` directory contains configuration files for the server. On `config.go` file, you can find the configuration for the server, database, and the JWT secret key, using `viper` to Unmarshal yaml files.
3. `internal` directory contains all the services and internal packages for the server, including new database connection using `GORM` for `PostgreSQL` and database entity `models`.
4. `tests` directory contains utilities required to run unit tests.
5. Each functional components on `internal` contains its own `repository`, which handles queries and database operations, `service` which handles business logic, and `grpc` which handles the gRPC methods.
<br/>
<br/>

## How to Run
1. Clone this repository
2. Install `docker` and `docker-compose` if not yet on your machine.
3. Run `docker compose -f docker-compose-test.yaml up` on the root directory of this repository.
4. Server will be available on `localhost:55000`
5. Feel free to test the API from Postman or other gRPC clients.

## Perform tests
```go
go test ./... --cover
```