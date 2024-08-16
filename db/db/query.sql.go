// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const createNote = `-- name: CreateNote :execresult
INSERT INTO notes (
    note_id,
    title,
    tags,
    content
) VALUES (?, ?, ?, ?)
`

type CreateNoteParams struct {
	NoteID  string
	Title   string
	Tags    string
	Content string
}

func (q *Queries) CreateNote(ctx context.Context, arg CreateNoteParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createNote,
		arg.NoteID,
		arg.Title,
		arg.Tags,
		arg.Content,
	)
}

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (
    user_id,
    name,
    email,
    address,
    sex,
    birthday,
    password
) VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreateUserParams struct {
	UserID   string
	Name     string
	Email    string
	Address  string
	Sex      int32
	Birthday string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.UserID,
		arg.Name,
		arg.Email,
		arg.Address,
		arg.Sex,
		arg.Birthday,
		arg.Password,
	)
}

const deleteNote = `-- name: DeleteNote :exec
DELETE FROM notes
WHERE note_id = ?
`

func (q *Queries) DeleteNote(ctx context.Context, noteID string) error {
	_, err := q.db.ExecContext(ctx, deleteNote, noteID)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, userID string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, userID)
	return err
}

const getNote = `-- name: GetNote :one
SELECT id, notes.note_id, title, tags, content, created_at, updated_at, user_id, user_notes.note_id FROM notes
JOIN user_notes ON notes.note_id = user_notes.note_id
WHERE user_notes.note_id = ? LIMIT 1
`

type GetNoteRow struct {
	ID        int64
	NoteID    string
	Title     string
	Tags      string
	Content   string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	UserID    string
	NoteID_2  string
}

func (q *Queries) GetNote(ctx context.Context, noteID string) (GetNoteRow, error) {
	row := q.db.QueryRowContext(ctx, getNote, noteID)
	var i GetNoteRow
	err := row.Scan(
		&i.ID,
		&i.NoteID,
		&i.Title,
		&i.Tags,
		&i.Content,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.NoteID_2,
	)
	return i, err
}

const getNotes = `-- name: GetNotes :many
SELECT id, notes.note_id, title, tags, content, created_at, updated_at, user_id, user_notes.note_id FROM notes
JOIN user_notes ON notes.note_id = user_notes.note_id
WHERE user_notes.user_id = ?
ORDER BY created_at DESC
`

type GetNotesRow struct {
	ID        int64
	NoteID    string
	Title     string
	Tags      string
	Content   string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	UserID    string
	NoteID_2  string
}

func (q *Queries) GetNotes(ctx context.Context, userID string) ([]GetNotesRow, error) {
	rows, err := q.db.QueryContext(ctx, getNotes, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetNotesRow
	for rows.Next() {
		var i GetNotesRow
		if err := rows.Scan(
			&i.ID,
			&i.NoteID,
			&i.Title,
			&i.Tags,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.NoteID_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, user_id, name, address, email, password, sex, birthday, created_at, updated_at FROM users
WHERE user_id = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, userID string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, userID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Address,
		&i.Email,
		&i.Password,
		&i.Sex,
		&i.Birthday,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, user_id, name, address, email, password, sex, birthday, created_at, updated_at FROM users
ORDER BY created_at DESC
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Address,
			&i.Email,
			&i.Password,
			&i.Sex,
			&i.Birthday,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateNote = `-- name: UpdateNote :exec
UPDATE notes
SET title = ?,
    tags = ?,
    content = ?
WHERE note_id = ?
`

type UpdateNoteParams struct {
	Title   string
	Tags    string
	Content string
	NoteID  string
}

func (q *Queries) UpdateNote(ctx context.Context, arg UpdateNoteParams) error {
	_, err := q.db.ExecContext(ctx, updateNote,
		arg.Title,
		arg.Tags,
		arg.Content,
		arg.NoteID,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET name = ?,
    email = ?,
    address = ?,
    sex = ?,
    birthday = ?,
    password = ?
WHERE user_id = ?
`

type UpdateUserParams struct {
	Name     string
	Email    string
	Address  string
	Sex      int32
	Birthday string
	Password string
	UserID   string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.Address,
		arg.Sex,
		arg.Birthday,
		arg.Password,
		arg.UserID,
	)
	return err
}
