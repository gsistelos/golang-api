// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

type Post struct {
	ID         string
	Content    string
	Visibility string
	UserID     string
}

type User struct {
	ID       string
	Username string
	Email    string
	Password string
}
