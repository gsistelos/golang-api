package server

import (
	"context"
	"sync"

	"github.com/google/uuid"
	v1 "github.com/gsistelos/golang-gRPC-API/gen/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	mu    sync.Mutex
	users map[string]*v1.User
	v1.UnimplementedUserServiceServer
}

func New() *Server {
	return &Server{
		users: make(map[string]*v1.User),
	}
}

func (s *Server) AddUser(ctx context.Context, req *v1.AddUserRequest) (*v1.AddUserResponse, error) {
	u := v1.User{
		Id:       uuid.New().String(),
		Username: req.Username,
		Email:    req.Email,
	}

	s.mu.Lock()
	s.users[u.Id] = &u
	s.mu.Unlock()

	return &v1.AddUserResponse{User: &u}, nil
}

func (s *Server) ListUsers(ctx context.Context, req *v1.ListUsersRequest) (*v1.ListUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	us := make([]*v1.User, req.Offset, req.Limit)
	for id := range s.users {
		u := s.users[id]
		us = append(us, u)
	}

	return &v1.ListUsersResponse{Users: us, Count: int32(len(us)), HasNext: len(s.users)-int(req.Offset) > int(req.Limit)}, nil
}

func (s *Server) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	u, ok := s.users[req.UserId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user with ID: %q does not exists", req.UserId)
	}

	return &v1.GetUserResponse{User: u}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	u, ok := s.users[req.User.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user with ID: %q does not exists", req.User.Id)
	}

	u.Username = req.User.Username
	u.Email = req.User.Email

	return &v1.UpdateUserResponse{User: u}, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	u, ok := s.users[req.UserId]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user with ID: %q does not exists", req.UserId)
	}
	delete(s.users, req.UserId)

	return &v1.DeleteUserResponse{UserId: u.Id}, nil
}
