-- name: InsertItemtMask :one
INSERT INTO item_masks (
    item_code,
    mask,
    mask_hash,
    created_by,
    created_at,
    item_id
)
VALUES ($1, $2, $3, $4, NOW(), $5)
RETURNING id, item_code, mask, mask_hash, created_by, created_at, item_id;

-- name: InsertItemMaskAnswer :exec
INSERT INTO item_mask_answers (mask_id, question_id, option_id, position)
VALUES ($1, $2, $3, $4);

-- SQLC query
-- name: GetOptionValueByID :one
SELECT value FROM question_options WHERE id = $1;
