-- name: CreateUser :exec
INSERT INTO users (
    id,
    name,
    email,
    password,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, now(), now()
);

-- name: GetUserByEmail :one
SELECT
    id,
    name,
    email,
    password,
    created_at,
    updated_at
FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT
    id,
    name,
    email,
    created_at,
    updated_at
FROM users
WHERE id = $1;
