-- name: CreateModifier :one
INSERT INTO modifier (
    description,
    created_by
) VALUES (
$1, $2
)
RETURNING *;
