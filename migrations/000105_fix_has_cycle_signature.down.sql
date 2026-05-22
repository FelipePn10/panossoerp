BEGIN;

-- Restore the legacy has_cycle(varchar, varchar) if rolling back.
-- Note: this function is broken on the current schema (bigint columns)
-- and should not be relied upon.
CREATE OR REPLACE FUNCTION has_cycle(start_code VARCHAR, target_code VARCHAR)
RETURNS boolean AS $$
WITH RECURSIVE ancestors AS (
    SELECT parent_code
    FROM item_structures
    WHERE child_code = start_code::BIGINT

    UNION ALL

    SELECT s.parent_code
    FROM item_structures s
    JOIN ancestors a ON s.child_code = a.parent_code
)
SELECT EXISTS (
    SELECT 1 FROM ancestors WHERE parent_code = target_code::BIGINT
);
$$ LANGUAGE sql;

COMMIT;
