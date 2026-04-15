-- name: CreateStructureComponent :one
INSERT INTO item_structures (
    parent_item_id,
    parent_code,
    child_item_id,
    child_code,
    parent_mask,
    quantity,
    unit_of_measurement,
    loss_percentage,
    position,
    notes,
    created_by
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- name: GetStructureComponentByID :one
SELECT *
FROM item_structures
WHERE id = $1;

-- name: GetAllDirectChildren :many
-- Retorna TODOS os filhos ativos de um pai (genéricos + mascarados). (999)
SELECT *
FROM item_structures
WHERE parent_item_id = $1
  AND is_active = TRUE
ORDER BY position, id;

-- name: GetGenericChildren :many
-- Retorna apenas os filhos GENÉRICOS (sem máscara) de um pai.
SELECT *
FROM item_structures
WHERE parent_item_id = $1
  AND parent_mask IS NULL
  AND is_active = TRUE
ORDER BY position, id;

-- name: GetDirectChildrenForMask :many
-- Retorna filhos ativos de um pai para uma máscara específica E genéricos.
-- A lógica de prioridade (específico > genérico) é aplicada na camada de aplicação.
SELECT *
FROM item_structures
WHERE parent_item_id = $1
  AND is_active = TRUE
  AND (parent_mask = $2 OR parent_mask IS NULL)
ORDER BY
    -- Específicos primeiro, para facilitar a deduplicação na app layer
    CASE WHEN parent_mask IS NOT NULL THEN 0 ELSE 1 END,
    position,
    id;

-- name: UpdateStructureComponent :one
UPDATE item_structures
SET
    quantity            = $2,
    unit_of_measurement = $3,
    loss_percentage     = $4,
    position            = $5,
    notes               = $6,
    updated_at          = NOW()
WHERE id = $1
  AND is_active = TRUE
RETURNING *;

-- name: DeactivateStructureComponent :exec
UPDATE item_structures
SET
    is_active  = FALSE,
    updated_at = NOW()
WHERE id = $1;

-- name: GetItemCodeAndDescription :one
-- Busca código e descrição técnica de um item pelo ID.
SELECT
    i.code::text AS code,
    COALESCE(
            p.pdm_description_technique,
            i.code::text
    ) AS description
FROM items i
         LEFT JOIN item_structures s
                   ON s.child_item_id = i.id
                       AND s.is_active = true
         LEFT JOIN items p
                   ON p.id = s.parent_item_id
WHERE i.id = $1
ORDER BY s.position ASC
    LIMIT 1;

-- name: ItemExists :one
SELECT EXISTS (
    SELECT 1
    FROM items
    WHERE id = $1
) AS "exists";

-- name: HasCyclicReference :one
-- Verifica se adicionar child_item_id como filho de parent_item_id
-- criaria uma referência circular na árvore BOM.
-- Retorna TRUE se houver ciclo (operação deve ser BLOQUEADA).
WITH RECURSIVE ancestors(ancestor_id) AS (
    SELECT s.parent_item_id
    FROM item_structures s
    WHERE s.child_item_id = $1::bigint
    AND s.is_active = TRUE

UNION ALL

SELECT s.parent_item_id
FROM item_structures s
         JOIN ancestors ON s.child_item_id = ancestors.ancestor_id
WHERE s.is_active = TRUE
    )
SELECT EXISTS (
    SELECT 1
    FROM ancestors
    WHERE ancestor_id = $2::bigint
) AS has_cycle;

-- name: GetItemMaskAnswersByValue :many
-- Retorna as respostas de uma máscara específica de um item,
-- usado para propagar a máscara do pai para os filhos.
SELECT
    ima.question_id,
    ima.position,
    ima.option_id
FROM item_masks im
JOIN item_mask_answers ima ON ima.mask_id = im.id
WHERE im.item_id    = $1   -- item pai
  AND im.mask = $2   -- ex: '100#100#50'
ORDER BY ima.position;

-- name: GetItemQuestions :many
-- Retorna as perguntas associadas a um item, ordenadas por posição.
-- Usado para calcular a máscara do filho com base nas respostas do pai.
SELECT
    iq.question_id,
    iq.position
FROM item_questions iq
WHERE iq.item_id = $1
ORDER BY iq.position;
