-- name: CreateUser :one
INSERT INTO users (
  username,
  password_hashed,
  fullname
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
  password_hashed = COALESCE(sqlc.narg(password_hashed), password_hashed),
  password_updated_at = COALESCE(sqlc.narg(password_updated_at), password_updated_at),
  fullname = COALESCE(sqlc.narg(fullname), fullname)
WHERE
  username = sqlc.arg(username)
RETURNING *;