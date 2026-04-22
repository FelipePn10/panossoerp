ALTER TABLE item_structures
    RENAME COLUMN position TO sequence;

WITH ranked AS (
    SELECT
        id,
        ROW_NUMBER() OVER (
            PARTITION BY parent_code
            ORDER BY sequence, id
        ) AS new_sequence
    FROM item_structures
    WHERE is_active = TRUE
)
UPDATE item_structures i
SET sequence = r.new_sequence
    FROM ranked r
WHERE i.id = r.id;

CREATE UNIQUE INDEX uq_structure_sequence
    ON item_structures (parent_code, sequence)
    WHERE is_active = TRUE;