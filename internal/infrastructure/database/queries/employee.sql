-- name: CreateEmployee :one
INSERT INTO employee (
    enterprise_id,
    code,
    description,
    name
) VALUES (
$1, $2, $3, $4
)
RETURNING *;
