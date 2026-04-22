DROP FUNCTION IF EXISTS has_cycle(BIGINT, BIGINT);

CREATE OR REPLACE FUNCTION has_cycle(start_code BIGINT, target_code BIGINT)
RETURNS boolean AS $$
WITH RECURSIVE descendants AS (
    SELECT child_code
    FROM item_structures
    WHERE parent_code = start_code
      AND is_active = TRUE

    UNION ALL

    SELECT s.child_code
    FROM item_structures s
    JOIN descendants d ON s.parent_code = d.child_code
    WHERE s.is_active = TRUE
)
SELECT EXISTS (
    SELECT 1 FROM descendants WHERE child_code = target_code
);
$$ LANGUAGE sql;