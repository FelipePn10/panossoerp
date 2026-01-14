-- name: CreateBom :one
INSERT INTO boms (
    product_id,
    mask,
    bom_type,
    version,
    status,
    valid_from
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;
