-- name: CreateEnterprise :one
INSERT INTO enterprise (
    id,
    code,
    name,
    created_by,
    created_at
) VALUES (
    $1, $2, $3, $4, now()
) RETURNING *;
