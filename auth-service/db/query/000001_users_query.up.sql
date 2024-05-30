-- name: CreateUser :one
INSERT INTO users(
    id,  username, email,  password, created_at, updated_at
) VALUES (
          $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUser :many
SELECT * FROM users;

-- name: UpdateUser :one
UPDATE users SET verified_email = $2, updated_at = $3 WHERE email = $1 RETURNING *;


-- name: DeleteUserByEmail :one
DELETE FROM users WHERE email = $1 RETURNING *;






