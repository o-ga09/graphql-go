-- name: GetNote :one
SELECT * FROM notes
WHERE note_id = ? LIMIT 1;

-- name: GetNotes :many
SELECT * FROM notes
ORDER BY created_at DESC;

-- name: CreateNote :execresult
INSERT INTO notes (
    note_id,
    title,
    content
) VALUES (?, ?, ?);

-- name: UpdateNote :exec
UPDATE notes
SET title = ?,
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
    password
) VALUES (?, ?, ?, ?);

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
    password = ?
WHERE user_id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = ?;
