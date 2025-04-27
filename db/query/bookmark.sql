-- name: CreateBook :one
INSERT INTO bookmark(
                     owner_id,
                     title,
                     url
) VALUES($1, $2, $3) RETURNING *;

-- name: GetBookMark :one
SELECT * FROM bookmark
WHERE id = $1
LIMIT 1;

-- name: DeleteBookMark :exec
DELETE FROM bookmark WHERE id = $1;

-- name: GetAllBookMarks :many
SELECT * FROM bookmark
LIMIT $1
OFFSET $2;

-- name: GetBookmarksByUser :many
SELECT * FROM bookmark
WHERE owner_id = $1
ORDER BY created_at DESC;