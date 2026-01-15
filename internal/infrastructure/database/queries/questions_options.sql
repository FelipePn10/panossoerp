-- name: CreateQuestionOption :one
INSERT INTO question_options (
    question_id,
    created_by,
    value
) VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetQuestionOptionByID :one
SELECT *
FROM question_options
WHERE id = $1;

-- name: DeleteQuestionOption :exec
DELETE FROM question_options
WHERE id = $1;

-- name: ExistsQuestionOptionByValue :one
SELECT *
FROM question_options
WHERE value = $1;