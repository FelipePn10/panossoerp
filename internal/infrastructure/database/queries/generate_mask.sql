-- name: InsertProductMask :one
INSERT INTO product_masks (
    product_code, 
    mask, 
    mask_hash, 
    created_by, 
    created_at
    )
VALUES ($1, $2, $3, $4, NOW())
RETURNING id, product_code, mask, mask_hash, created_by, created_at;

-- name: InsertProductMaskAnswer :exec
INSERT INTO product_mask_answers (mask_id, question_id, option_id, position)
VALUES ($1, $2, $3, $4);

-- SQLC query
-- name: GetOptionValueByID :one
SELECT value FROM question_options WHERE id = $1;

