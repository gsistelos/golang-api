-- name: CreateUser :execresult
INSERT INTO users (
    id, username, email, password
) VALUES (
    ?, ?, ?, ?
);

-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
LIMIT ? OFFSET ?;

-- name: UpdateUser :execresult
UPDATE users
SET username = ?, email = ?, password = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
