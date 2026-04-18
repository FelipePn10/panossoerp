ALTER TABLE item_structures
    ALTER COLUMN parent_code DROP NOT NULL;

ALTER TABLE item_structures
    ALTER COLUMN child_code DROP NOT NULL;