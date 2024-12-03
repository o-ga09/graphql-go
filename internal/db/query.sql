-- name: GetNote :one
SELECT id, notes.note_id, title, tags, content, created_at, updated_at, user_id FROM notes
JOIN user_notes ON notes.note_id = user_notes.note_id
WHERE user_notes.note_id = ? AND deleted_at IS NULL LIMIT 1;

-- name: GetNotes :many
SELECT id, notes.note_id, title, tags, content, created_at, updated_at, user_id FROM notes
JOIN user_notes ON notes.note_id = user_notes.note_id
WHERE user_notes.user_id = ? AND deleted_at IS NULL
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
UPDATE notes
SET deleted_at = CURRENT_TIMESTAMP
WHERE note_id = ?;

-- name: CreateUser :execresult
INSERT INTO users (
    user_id,
    username,
    displayname
) VALUES (?, ?, ?);

-- name: GetUser :one
SELECT id, user_id, username, displayname, created_at, updated_at FROM users
WHERE user_id = ? AND deleted_at IS NULL LIMIT 1;

-- name: GetUsers :many
SELECT id, user_id, username, displayname, created_at, updated_at FROM users
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateUser :exec
UPDATE users
SET username = ?,
    displayname = ?    
WHERE user_id = ?;

-- name: DeleteUser :exec
UPDATE users
SET deleted_at = CURRENT_TIMESTAMP
WHERE user_id = ?;

-- name: CreateUserNote :exec
INSERT INTO user_notes (
    user_id,
    note_id
) VALUES (?, ?);
