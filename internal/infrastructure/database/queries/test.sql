-- name: CreateTest :one
INSERT INTO tests DEFAULT VALUES
RETURNING id, created_at;

-- name: GetTestByID :one
SELECT id, created_at
FROM tests
WHERE id = $1;

-- name: ListTests :many
SELECT id, created_at
FROM tests
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;
