-- name: GetProductMaskByItemCode :one
SELECT *
FROM item_masks
WHERE item_code = $1
ORDER BY created_at DESC
LIMIT 1;

-- name: DeleteItemMask :exec
DELETE FROM item_masks
WHERE id = $1;
