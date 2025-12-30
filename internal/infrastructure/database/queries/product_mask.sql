-- name: GetProductMaskByProductID :one
SELECT *
FROM product_masks
WHERE product_id = $1
ORDER BY created_at DESC
LIMIT 1;

-- name: DeleteProductMask :exec
DELETE FROM product_masks
WHERE id = $1;