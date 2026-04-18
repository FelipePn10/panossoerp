CREATE OR REPLACE FUNCTION has_cycle(start_code BIGINT, target_code BIGINT)
RETURNS boolean AS $$
WITH RECURSIVE ancestors AS (
    SELECT parent_code
    FROM item_structures
    WHERE child_code = start_code

    UNION ALL

    SELECT s.parent_code
    FROM item_structures s
    JOIN ancestors a ON s.child_code = a.parent_code
)
SELECT EXISTS (
    SELECT 1 FROM ancestors WHERE parent_code = target_code
);
$$ LANGUAGE sql;