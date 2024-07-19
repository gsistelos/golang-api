package server

import (
	"context"
	"database/sql"

	database "github.com/gsistelos/grpc-api/gen/db"
	v1 "github.com/gsistelos/grpc-api/gen/user/v1"
)

type Server struct {
	ctx     context.Context
	queries *database.Queries
	v1.UnimplementedUserServiceServer
}

func New(db *sql.DB) *Server {
	return &Server{
		ctx:     context.Background(),
		queries: database.New(db),
	}
}
