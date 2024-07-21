-- name: CreatePost :execresult
INSERT INTO posts (
    id, content, visibility, user_id
) VALUES (
    ?, ?, ?, ?
);

-- name: GetPost :one
SELECT * FROM posts
WHERE id = ? LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
LIMIT ? OFFSET ?;

-- name: ListPostsByUser :many
SELECT * FROM posts
WHERE user_id = ?
LIMIT ? OFFSET ?;

-- name: UpdatePost :execresult
UPDATE posts
SET content = ?, visibility = ?
WHERE id = ?;

-- name: DeletePost :exec
DELETE FROM posts 
WHERE id = ?;
