-- name: CreateUsers :one
INSERT INTO users(
    id, full_name, username, email, photo, user_active, password, created_at, updated_at
) VALUES (
          $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUser :many
SELECT * FROM users;
