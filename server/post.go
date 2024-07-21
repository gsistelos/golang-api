package server

import (
	"context"

	"github.com/google/uuid"
	v1 "github.com/gsistelos/golang-api/gen/post/v1"
	sqlc "github.com/gsistelos/golang-api/gen/sqlc"
)

func PostSqlcToV1(p *sqlc.Post) *v1.Post {
	return &v1.Post{
		Id:         p.ID,
		Content:    p.Content,
		Visibility: p.Visibility,
		UserId:     p.UserID,
	}
}

func (s *Server) AddPost(ctx context.Context, req *v1.AddPostRequest) (*v1.AddPostResponse, error) {
	id := uuid.NewString()

	_, err := s.queries.CreatePost(ctx, sqlc.CreatePostParams{
		ID:         id,
		Content:    req.Content,
		Visibility: req.Visibility,
		UserID:     req.UserId,
	})
	if err != nil {
		return nil, err
	}

	post, err := s.queries.GetPost(ctx, id)
	if err != nil {
		return nil, err
	}

	return &v1.AddPostResponse{Post: PostSqlcToV1(&post)}, nil
}

func (s *Server) GetPost(ctx context.Context, req *v1.GetPostRequest) (*v1.GetPostResponse, error) {
	post, err := s.queries.GetPost(ctx, req.PostId)
	if err != nil {
		return nil, err
	}

	return &v1.GetPostResponse{Post: PostSqlcToV1(&post)}, nil
}

func (s *Server) ListPosts(ctx context.Context, req *v1.ListPostsRequest) (*v1.ListPostsResponse, error) {
	posts, err := s.queries.ListPosts(ctx, sqlc.ListPostsParams{Limit: req.Limit, Offset: req.Offset})
	if err != nil {
		return nil, err
	}

	v1Posts := make([]*v1.Post, len(posts))
	for i, post := range posts {
		v1Posts[i] = PostSqlcToV1(&post)
	}

	return &v1.ListPostsResponse{Posts: v1Posts}, nil
}

func (s *Server) ListPostsByUser(ctx context.Context, req *v1.ListPostsByUserRequest) (*v1.ListPostsByUserResponse, error) {
	posts, err := s.queries.ListPostsByUser(ctx, sqlc.ListPostsByUserParams{UserID: req.UserId, Limit: req.Limit, Offset: req.Offset})
	if err != nil {
		return nil, err
	}

	v1Posts := make([]*v1.Post, len(posts))
	for i, post := range posts {
		v1Posts[i] = PostSqlcToV1(&post)
	}

	return &v1.ListPostsByUserResponse{Posts: v1Posts}, nil
}

func (s *Server) UpdatePost(ctx context.Context, req *v1.UpdatePostRequest) (*v1.UpdatePostResponse, error) {
	_, err := s.queries.UpdatePost(ctx, sqlc.UpdatePostParams{
		ID:         req.PostId,
		Content:    req.Content,
		Visibility: req.Visibility,
	})
	if err != nil {
		return nil, err
	}

	post, err := s.queries.GetPost(ctx, req.PostId)
	if err != nil {
		return nil, err
	}

	return &v1.UpdatePostResponse{Post: PostSqlcToV1(&post)}, nil
}

func (s *Server) DeletePost(ctx context.Context, req *v1.DeletePostRequest) (*v1.DeletePostResponse, error) {
	err := s.queries.DeletePost(ctx, req.PostId)
	if err != nil {
		return nil, err
	}

	return &v1.DeletePostResponse{}, nil
}
