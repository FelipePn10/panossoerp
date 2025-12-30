-- name: CreateQuestion :one 
INSERT INTO questions (
    name,
    createdby
) VALUES (
    $1,
    $2
) RETURNING *;
