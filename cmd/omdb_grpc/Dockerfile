FROM golang:1.20
WORKDIR /app
COPY . .
RUN go build -o server ./cmd/omdb_grpc
CMD ["/app/server"]
