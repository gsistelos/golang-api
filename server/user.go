package server

import (
	"context"

	"github.com/google/uuid"
	sqlc "github.com/gsistelos/golang-api/gen/sqlc"
	v1 "github.com/gsistelos/golang-api/gen/user/v1"
	"golang.org/x/crypto/bcrypt"
)

func UserSqlcToV1(u *sqlc.User) *v1.User {
	return &v1.User{
		Id:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}

func (s *Server) AddUser(ctx context.Context, req *v1.AddUserRequest) (*v1.AddUserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	id := uuid.NewString()

	_, err = s.queries.CreateUser(ctx, sqlc.CreateUserParams{
		ID:       id,
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return nil, err
	}

	user, err := s.queries.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &v1.AddUserResponse{User: UserSqlcToV1(&user)}, nil
}

func (s *Server) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	user, err := s.queries.GetUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &v1.GetUserResponse{User: UserSqlcToV1(&user)}, nil
}

func (s *Server) ListUsers(ctx context.Context, req *v1.ListUsersRequest) (*v1.ListUsersResponse, error) {
	users, err := s.queries.ListUsers(ctx, sqlc.ListUsersParams{Limit: req.Limit, Offset: req.Offset})
	if err != nil {
		return nil, err
	}

	v1Users := make([]*v1.User, len(users))
	for i, user := range users {
		v1Users[i] = UserSqlcToV1(&user)
	}

	return &v1.ListUsersResponse{Users: v1Users}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	_, err = s.queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:       req.UserId,
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return nil, err
	}

	user, err := s.queries.GetUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateUserResponse{User: UserSqlcToV1(&user)}, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	err := s.queries.DeleteUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteUserResponse{}, nil
}
