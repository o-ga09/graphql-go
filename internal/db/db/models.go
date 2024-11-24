// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
)

type Note struct {
	ID        int64
	NoteID    string
	Title     string
	Tags      string
	Content   string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type User struct {
	ID        int64
	UserID    string
	Name      string
	Address   string
	Email     string
	Password  string
	Sex       int32
	Birthday  string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type UserNote struct {
	UserID string
	NoteID string
}
