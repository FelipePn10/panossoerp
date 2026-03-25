-- name: CreateGroup :one
INSERT INTO groups (
    code,
    description,
    enterprise_id,
    created_by
) VALUES (
$1, $2, $3, $4
)
RETURNING *;
