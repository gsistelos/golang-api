package server

import (
	"sync"

	v1 "github.com/gsistelos/grpc-api/gen/user/v1"
)

type Server struct {
	mu sync.Mutex
	us map[string]*v1.User
	v1.UnimplementedUserServiceServer
}

func New() *Server {
	return &Server{
		us: make(map[string]*v1.User),
	}
}
