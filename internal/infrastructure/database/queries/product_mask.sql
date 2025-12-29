-- name: CreateProductMask :one
INSERT INTO product_masks (
    id,
    product_id,
    mask,
    mask_hash,
    business_id,
    created_by,
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

-- name: DeleteProductMask :exec
DELETE FROM product_masks
WHERE id = $1;