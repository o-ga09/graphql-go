// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

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
	DeleteAt  sql.NullTime
}

type User struct {
	ID          int64
	UserID      string
	Username    string
	Displayname string
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	DeleteAt    sql.NullTime
}

type UserNote struct {
	UserID string
	NoteID string
}
