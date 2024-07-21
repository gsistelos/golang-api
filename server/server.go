package server

import (
	"context"
	"database/sql"

	postV1 "github.com/gsistelos/golang-api/gen/post/v1"
	sqlc "github.com/gsistelos/golang-api/gen/sqlc"
	userV1 "github.com/gsistelos/golang-api/gen/user/v1"
)

type Server struct {
	ctx     context.Context
	queries *sqlc.Queries
	postV1.UnimplementedPostServiceServer
	userV1.UnimplementedUserServiceServer
}

func New(db *sql.DB) *Server {
	return &Server{
		ctx:     context.Background(),
		queries: sqlc.New(db),
	}
}
