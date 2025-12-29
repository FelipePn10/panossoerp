-- name: CreateProduct :one
INSERT INTO products (
    id,
    code,
    group_code,
    name,
    created_by,
    created_at,
    updated_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    NOW(),
    NOW()
)
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;