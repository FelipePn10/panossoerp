-- name: CreateQuestion :one
INSERT INTO questions (
    name,
    createdby
) VALUES (
    $1,
    $2
) RETURNING *;

-- name: GetQuestionByID :one
SELECT *
FROM questions
WHERE id = $1;

-- name: DeleteQuestion :exec
DELETE FROM questions
WHERE id = $1;

-- name: FindQuestionByName :many
SELECT *
FROM questions
WHERE name = $1;

-- name: ExistsQuestionByName :one
SELECT EXISTS (
    SELECT 1 FROM questions WHERE name = $1
);
