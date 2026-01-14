-- name: CreateProduct :one
INSERT INTO products (
    id,
    code,
    group_code,
    name,
    created_by,
    created_at
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    NOW()
)
RETURNING *;

-- name: FindByNameAndCode :one
SELECT *
FROM products
WHERE name = $1 AND code = $2;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;