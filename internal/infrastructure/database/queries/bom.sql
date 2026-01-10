-- name: CreateBom :one
INSERT INTO boms (
    id,
    product_id,
    bom_type,
    version,
    status,
    valid_from,
    created_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    NOW()
)
RETURNING *;
