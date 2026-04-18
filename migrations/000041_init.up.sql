UPDATE item_structures
SET parent_code = ''
WHERE parent_code IS NULL;

UPDATE item_structures
SET child_code = ''
WHERE child_code IS NULL;

ALTER TABLE item_structures
    ALTER COLUMN parent_code SET NOT NULL;

ALTER TABLE item_structures
    ALTER COLUMN child_code SET NOT NULL;