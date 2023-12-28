-- name: CreateUser :one
INSERT INTO users(name, email, password)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: UpdateUser :exec
UPDATE users
SET name = $2, email = $3, password = $4
WHERE id = $1;
