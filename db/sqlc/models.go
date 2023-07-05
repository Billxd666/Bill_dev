// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import ()

type Account struct {
	ID    int64  `json:"id"`
	Owner string `json:"owner"`
	Email string `json:"email"`
}

type Post struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  int64  `json:"author"`
}

type User struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}
