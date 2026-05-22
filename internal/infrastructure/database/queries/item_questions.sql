-- name: AssociateQuestionItem :exec
INSERT INTO item_questions (
    item_code,
    question_id,
    position,
    created_at
) VALUES ($1, $2, $3, $4);

-- name: ExistsByItemAndQuestion :one
SELECT EXISTS (
    SELECT 1
    FROM item_questions
    WHERE item_code = $1
      AND question_id = $2
);

-- name: ExistsByItemAndPosition :one
SELECT EXISTS (
    SELECT 1
    FROM item_questions
    WHERE item_code = $1
      AND position = $2
);

-- name: GetQuestionsByItemCode :many
SELECT iq.item_code, iq.question_id, iq.position, iq.created_at, q.name AS question_name
FROM item_questions iq
JOIN questions q ON q.id = iq.question_id
JOIN items i ON i.id = iq.item_code
WHERE i.code = $1
ORDER BY iq.position;

-- name: ListAllItemQuestions :many
SELECT iq.item_code, iq.question_id, iq.position, iq.created_at,
       i.code AS item_business_code, q.name AS question_name
FROM item_questions iq
JOIN questions q ON q.id = iq.question_id
JOIN items i ON i.id = iq.item_code
ORDER BY i.code, iq.position;
