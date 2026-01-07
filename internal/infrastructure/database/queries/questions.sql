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

-- name: FindQuestionByNameAndCode :one
SELECT *
FROM questions
WHERE name = $1;