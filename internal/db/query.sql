-- name: GetNote :one
SELECT id, notes.note_id, title, tags, content, created_at, updated_at, user_id FROM notes
JOIN user_notes ON notes.note_id = user_notes.note_id
WHERE user_notes.note_id = ? LIMIT 1;

-- name: GetNotes :many
SELECT id, notes.note_id, title, tags, content, created_at, updated_at, user_id FROM notes
JOIN user_notes ON notes.note_id = user_notes.note_id
WHERE user_notes.user_id = ?
ORDER BY created_at DESC;

-- name: CreateNote :execresult
INSERT INTO notes (
    note_id,
    title,
    tags,
    content
) VALUES (?, ?, ?, ?);

-- name: UpdateNote :exec
UPDATE notes
SET title = ?,
    tags = ?,
    content = ?
WHERE note_id = ?;

-- name: DeleteNote :exec
DELETE FROM notes
WHERE note_id = ?;

-- name: CreateUser :execresult
INSERT INTO users (
    user_id,
    name,
    email,
    address,
    sex,
    birthday,
    password
) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetUser :one
SELECT * FROM users
WHERE user_id = ? LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY created_at DESC;

-- name: UpdateUser :exec
UPDATE users
SET name = ?,
    email = ?,
    address = ?,
    sex = ?,
    birthday = ?,
    password = ?
WHERE user_id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = ?;
