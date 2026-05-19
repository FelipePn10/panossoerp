-- name: GetProductMaskByItemCode :one
SELECT *
FROM item_masks
WHERE item_code = $1
ORDER BY created_at DESC
LIMIT 1;

-- name: DeleteItemMask :exec
DELETE FROM item_masks
WHERE id = $1;

-- name: ListAllItemMasks :many
SELECT id, item_code, mask, mask_hash, created_by, created_at
FROM item_masks
ORDER BY item_code, created_at DESC;
