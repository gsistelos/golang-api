package server

import (
	"context"
	"database/sql"

	sqlc "github.com/gsistelos/grpc-api/gen/sqlc"
	v1 "github.com/gsistelos/grpc-api/gen/user/v1"
)

type Server struct {
	ctx     context.Context
	queries *sqlc.Queries
	v1.UnimplementedUserServiceServer
}

func New(db *sql.DB) *Server {
	return &Server{
		ctx:     context.Background(),
		queries: sqlc.New(db),
	}
}
