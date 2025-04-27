-- name: CreateUserAccount :one
INSERT INTO userprofile (
                  username,
                  email,
                  hashed_password
) VALUES ($1, $2, $3) RETURNING *;

-- name: GetUserAccount :one
SELECT * FROM userprofile
WHERE id = $1 LIMIT 1;

-- name: GetAllUserAccounts :many
SELECT * FROM userprofile
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteUserAccount :exec
DELETE FROM userprofile
       WHERE id = $1;