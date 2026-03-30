-- name: CreateEmployee :one
INSERT INTO groups (
    enterprise_id,
    code,
    name,
    description
) VALUES (
$1, $2, $3, $4
)
RETURNING *;
