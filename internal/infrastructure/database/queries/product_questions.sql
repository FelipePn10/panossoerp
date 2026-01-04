-- name: AssociateQuestionProduct :exec
INSERT INTO product_questions (
    product_id,
    question_id,
    position,
    created_at
) VALUES ($1, $2, $3, $4);

-- name: ExistsByProductAndQuestion :one
SELECT EXISTS (
    SELECT 1
    FROM product_questions
    WHERE product_id = $1
      AND question_id = $2
);

-- name: ExistsByProductAndPosition :one
SELECT EXISTS (
    SELECT 1
    FROM product_questions
    WHERE product_id = $1
      AND position = $2
);
