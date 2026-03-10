-- name: CreateItem :one
INSERT INTO items (
    id,
    warehouse_id
    code,
    name,
    desc,
    type,
    status,
    health,
    created_by,
    created_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
    NOW()
)
RETURNING *;